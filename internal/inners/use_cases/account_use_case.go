package use_cases

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"social-media-backend-1/internal/inners/models/entities"
	"social-media-backend-1/internal/outers/repositories"
	"strings"
)

type AccountUseCase struct {
	AccountRepository *repositories.AccountRepository
	FileRepository    *repositories.FileRepository
}

func NewAccountUseCase(accountRepository *repositories.AccountRepository, fileRepository *repositories.FileRepository) *AccountUseCase {
	return &AccountUseCase{
		AccountRepository: accountRepository,
		FileRepository:    fileRepository,
	}
}

func (uc *AccountUseCase) GetAllAccounts(ctx context.Context) ([]*entities.Account, error) {
	foundAccounts, err := uc.AccountRepository.GetAllAccounts(ctx)
	if err != nil {
		return nil, err
	}

	for _, foundAccount := range foundAccounts {
		foundAccount.ImageURL, err = uc.GetAccountImageURL(ctx, foundAccount)
		if err != nil {
			return nil, err
		}
	}

	return foundAccounts, nil
}

func (uc *AccountUseCase) GetAccountsByIDs(ctx context.Context, ids []*uuid.UUID) ([]*entities.Account, error) {
	foundAccounts, err := uc.AccountRepository.GetAccountsByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	for _, foundAccount := range foundAccounts {
		foundAccount.ImageURL, err = uc.GetAccountImageURL(ctx, foundAccount)
		if err != nil {
			return nil, err
		}
	}

	return foundAccounts, nil
}

func (uc *AccountUseCase) GetAccountByID(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
	foundAccount, err := uc.AccountRepository.GetAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}

	foundAccount.ImageURL, err = uc.GetAccountImageURL(ctx, foundAccount)
	if err != nil {
		return nil, err
	}

	return foundAccount, nil
}

func (uc *AccountUseCase) CreateAccount(ctx context.Context, accountToCreate *entities.Account) (*entities.Account, error) {
	createdAccount, err := uc.AccountRepository.CreateAccount(ctx, accountToCreate)
	if err != nil {
		return nil, err
	}

	err = uc.UploadAccountImage(ctx, accountToCreate)
	if err != nil {
		return nil, err
	}

	createdAccount.ImageURL, err = uc.GetAccountImageURL(ctx, createdAccount)
	if err != nil {
		return nil, err
	}

	return createdAccount, nil
}

func (uc *AccountUseCase) UpdateAccountByID(ctx context.Context, id uuid.UUID, accountToUpdate *entities.Account) (*entities.Account, error) {
	foundAccount, err := uc.AccountRepository.GetAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = uc.DeleteAccountImage(ctx, foundAccount)
	if err != nil {
		return nil, err
	}

	err = uc.UploadAccountImage(ctx, accountToUpdate)
	if err != nil {
		return nil, err
	}

	updatedAccount, err := uc.AccountRepository.UpdateAccountByID(ctx, id, accountToUpdate)
	if err != nil {
		return nil, err
	}

	updatedAccount.ImageURL, err = uc.GetAccountImageURL(ctx, updatedAccount)
	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}

func (uc *AccountUseCase) DeleteAccountByID(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
	deletedAccount, err := uc.AccountRepository.DeleteAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return deletedAccount, nil
}

func (uc *AccountUseCase) DeleteAccountImage(ctx context.Context, account *entities.Account) error {
	if account.ImageID == nil {
		return nil
	}

	bucketName := "social-media-backend.account"
	objectName := account.ImageID.String()

	err := uc.FileRepository.Delete(ctx, bucketName, objectName)
	if err != nil {
		return fmt.Errorf("failed to delete account image: %w", err)
	}

	return nil
}

func (uc *AccountUseCase) UploadAccountImage(ctx context.Context, account *entities.Account) error {
	if account.Image == nil {
		return nil
	}

	imageID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	account.ImageID = &imageID

	file := account.Image.File
	fileSize := account.Image.Size
	fileExtension := strings.SplitN(account.Image.Filename, ".", 2)[1]
	bucketName := "social-media-backend.account"
	objectName := imageID.String()
	extensionToContentType := map[string]string{
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"png":  "image/png",
	}
	fileContentType, exists := extensionToContentType[fileExtension]
	if !exists {
		return fmt.Errorf("unsupported file extension: %s", fileExtension)
	}

	err = uc.FileRepository.Upload(ctx, bucketName, objectName, file, fileSize, fileContentType)
	if err != nil {
		return err
	}

	return nil
}

func (uc *AccountUseCase) GetAccountImageURL(ctx context.Context, account *entities.Account) (*string, error) {
	if account.ImageID == nil {
		return nil, nil
	}

	bucketName := "social-media-backend.account"
	objectName := account.ImageID.String()

	imageURL, err := uc.FileRepository.GetURL(ctx, bucketName, objectName)
	if err != nil {
		return nil, err
	}

	return &imageURL, nil
}
