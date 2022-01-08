import * as React from 'react';
import { ChangeEvent, FormEvent } from 'react';
import { Button } from '@yandex/ui/Button/desktop/bundle';
import UserModel from '@models/UserModel';
import Validator from '@helpers/validator';
import { ContentCard } from '@components/ContentCard';
import { Textinput } from '@components/Textinput';
import { Info } from '@components/Info';
import { LoginViewProps, LoginViewState } from './LoginView.typings';
import { loginCn, submitCn } from './LoginView.consts';

import './index.scss';
import { Navigate } from 'react-router-dom';
import { ClientRoutes } from '@consts/routes';

export default class LoginView extends React.Component<LoginViewProps, LoginViewState> {
    private userModel: UserModel;
    private validator: Validator;

    constructor(props: LoginViewProps) {
        super(props);

        this.userModel = new UserModel();
        this.validator = new Validator();

        this.state = {
            login: '',
            password: '',
            progress: false,
            showAlert: false,
            isAuth: false,
        };
    }

    onChangeLogin = (event: ChangeEvent<HTMLInputElement>) => {
        this.setState({ login: event.target.value });
    };

    onChangePassword = (event: ChangeEvent<HTMLInputElement>) => {
        this.setState({ password: event.target.value });
    };

    onSubmit = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const {
            login,
            password,
        } = this.state;

        this.setState({
            progress: true,
            showAlert: false,
        });

        // TODO: Замедление срабатывания метода для демонстрации загрузки
        // await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 1000));

        if (!this.validator.validatePassword(password)) {
            this.setState({
                progress: false,
                showAlert: true,
                password: '',
            });
            return;
        }

        const response = await this.userModel.login(login, password);
        if (response.ok) {
            this.props.setUser && this.props.setUser( await this.userModel.getUserProfile() );
        } else {
            this.setState({
                progress: false,
                showAlert: true,
                password: '',
            });
        }
    };

    render() {
        const {
            login,
            password,
            showAlert,
            progress,
        } = this.state;

        if (this.props.user) {
            return <Navigate to={ClientRoutes.profilePage} replace={true} />;
        }

        return (
            <ContentCard className={loginCn}>
                <Info
                    show={showAlert}
                    type="alert"
                >
                    Неверный логин и/или пароль
                </Info>
                <form onSubmit={this.onSubmit}>
                    <Textinput
                        onChange={this.onChangeLogin}
                        value={login}
                        label="Логин"
                    />
                    <Textinput
                        onChange={this.onChangePassword}
                        value={password}
                        label="Пароль"
                        type="password"
                    />
                    <Button
                        className={submitCn}
                        view="action"
                        size="m"
                        type="submit"
                        progress={progress}
                    >
                        Войти
                    </Button>
                </form>
            </ContentCard>
        );
    }
}
