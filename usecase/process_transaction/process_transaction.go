package process_transaction

import "github.vom/jin-cloud-max/imersao5/entity"

type ProcessTransaction struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transation := entity.NewTransaction()

	transation.ID = input.ID
	transation.AccountId = input.AccountID
	transation.Amount = input.Amount
	invalidTransaction := transation.IsValid()

	if invalidTransaction != nil {
		err := p.Repository.Insert(transation.ID, transation.AccountId, transation.Amount, "rejected", invalidTransaction.Error())

		if err != nil {
			return TransactionDtoOutput{}, err
		}

		output := TransactionDtoOutput{
			ID:           transation.ID,
			Status:       "rejected",
			ErrorMessage: invalidTransaction.Error(),
		}

		return output, nil
	}

	err := p.Repository.Insert(transation.ID, transation.AccountId, transation.Amount, "approved", "")

	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transation.ID,
		Status:       "approved",
		ErrorMessage: "",
	}

	return output, nil
}
