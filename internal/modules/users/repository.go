package users
import "context"

type UserModuleInterface interface {
    ListUsers(ctx context.Context) ([]DBClient, error)
    FindUserById(ctx context.Context, uuid string) (*DBClient, error)
    CreateUser(ctx context.Context, body DBClient) (*DBClient, error)
    DeleteUser(ctx context.Context, uuid string) (*DBClient, error)
}
