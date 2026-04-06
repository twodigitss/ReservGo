package users
import "context"

type UserModuleInterface interface {
    FindUsers(ctx context.Context) (*DBClients, error)
}
