import * as React from 'react';
import { ChangeEvent, FormEvent } from 'react';
import { Textinput } from '@yandex/ui/Textinput/desktop/bundle';
import { Button } from '@yandex/ui/Button/desktop/bundle';
import { Text } from '@yandex/ui/Text/desktop/bundle';
import UserService from '@services/UserService';
import Validator from '@helpers/validator';
import { ContentCard } from '@components/ContentCard';
import { Info } from '@components/Info';
import { LoginViewProps, LoginViewState } from './LoginView.typings';
import { inputCn, labelCn, loginCn, submitCn } from './LoginView.consts';

import './index.scss';
import { Navigate } from 'react-router-dom';

export default class LoginView extends React.Component<LoginViewProps, LoginViewState> {
    private userService: UserService;
    private validator: Validator;

    constructor(props: LoginViewProps) {
        super(props);

        this.userService = new UserService();
        this.validator = new Validator();

        this.state = {
            login: '',
            password: '',
            progress: false,
            showAlert: false,
            user: false,
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

        // Замедление срабатывания метода для демонстрации загрузки
        await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 2000));

        if (!this.validator.validateLogin(login) || !this.validator.validatePassword(password)) {
            this.setState({
                progress: false,
                showAlert: true,
                password: '',
            });
            return;
        }
        this.setState({ user: await this.userService.login(login, password) });
    };

    render() {
        const {
            login,
            password,
            showAlert,
            progress,
            user,
        } = this.state;

        return (
            <ContentCard className={loginCn}>
                {user && (
                    <Navigate to="/profile" replace={true} />
                )}
                <Info
                    show={showAlert}
                    type="alert"
                >
                    Неверный логин и/или пароль
                </Info>
                <form onSubmit={this.onSubmit}>
                    <Text className={labelCn}>Логин</Text>
                    <Textinput
                        onChange={this.onChangeLogin}
                        value={login}
                        className={inputCn}
                        size="m"
                        view="default"
                        required
                    />
                    <Text className={labelCn}>Пароль</Text>
                    <Textinput
                        label="Пароль"
                        onChange={this.onChangePassword}
                        value={password}
                        className={inputCn}
                        size="m"
                        view="default"
                        type="password"
                        required
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
