package users
import "context"

//recordar que en mittal me ensenaron que es bueno tambien poner
//el nombre del storedProc como nombre del metodo porque asi te
//confundes menos

type UserService struct {
    repo UserModuleInterface
}

func NewService(repo UserModuleInterface) *UserService {
    return &UserService{repo: repo}
}

func (this *UserService) ListAllUsers(ctx context.Context) (*DBClients, error) {
	// delega al repo, no reimplementa
	// perfectamente puedo poner aqui logica que maneje el ouput 
	return this.repo.FindUsers(ctx)
}
