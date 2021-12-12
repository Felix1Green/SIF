import { useNavigate } from 'react-router-dom';
import { RoutesClientApi } from '../../consts/routes';

class UserService {
    navigate = useNavigate();

    async logIn(login?: string, password?: string) {
        if (!login || !password) {
            // eslint-disable-next-line no-console
            console.log('FAILURE');
            return 'FAILURE';
        }

        this.navigate(RoutesClientApi.profilePage);
        // eslint-disable-next-line no-console
        console.log('SUCCESS');
        return 'SUCCESS';
    }
}

export default UserService;
