package users
import "context"

type UserService struct {
    repo UserModuleInterface
}

func NewService(repo UserModuleInterface) *UserService {
    return &UserService{repo: repo}
}

func (this *UserService) ListAllUsers(ctx context.Context) ([]DBClient, error) {
	// delega al repo, no reimplementa
	return this.repo.ListUsers(ctx)
	// perfectamente puedo poner aqui logica que maneje el ouput
}

func (this *UserService) FindUserById(ctx context.Context, uuid string) (*DBClient, error){
	return this.repo.FindUserById(ctx, uuid)
}

func (this *UserService) CreateUser(ctx context.Context, body DBClient)(*DBClient, error){
	return this.repo.CreateUser(ctx, body)
}

func (this *UserService) DeleteUser(ctx context.Context, uuid string) (*DBClient, error){
	return this.repo.DeleteUser(ctx, uuid)
}
