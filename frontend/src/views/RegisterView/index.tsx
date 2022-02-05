import * as React from 'react';
import { RegisterViewProps, RegisterViewState } from '@views/RegisterView/RegisterView.typings';
import Validator from '@helpers/validator';
import UserModel from '@models/UserModel';
import { ChangeEvent, FormEvent } from 'react';
import { ClientRoutes } from '@consts/routes';
import { Navigate } from 'react-router-dom';
import { ContentCard } from '@src/components/ContentCard';
import { registrationCn, registrationSubmitCn } from './RegisterView.consts';
import { Info } from '@components/Info';
import { Textinput } from '@components/Textinput';
import { Button } from '@yandex/ui/Button/desktop/bundle';

import './index.scss';

export default class RegisterView extends React.Component<RegisterViewProps, RegisterViewState>{
    private userModel: UserModel;
    private validator: Validator;

    constructor(props: RegisterViewProps) {
        super(props);

        this.userModel = new UserModel();
        this.validator = new Validator();

        this.state = {
            login: '',
            name: '',
            surname: '',
            password: '',
            progress: false,
            showAlert: false,
            isSuccess: false,
        };
    }

    onChangeLogin = (event: ChangeEvent<HTMLInputElement>) => {
        this.setState({ login: event.target.value });
    };

    onChangePassword = (event: ChangeEvent<HTMLInputElement>) => {
        this.setState({ password: event.target.value });
    };

    onChangeName = (event: ChangeEvent<HTMLInputElement>) => {
        this.setState({ name: event.target.value });
    };

    onChangeSurname = (event: ChangeEvent<HTMLInputElement>) => {
        this.setState({ surname: event.target.value });
    };

    onSubmit = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const {
            login,
            password,
            name,
            surname,
        } = this.state;

        this.setState({
            progress: true,
            showAlert: false,
        });

        // TODO: Замедление срабатывания метода для демонстрации загрузки
        // await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 1000));

        if (
            !this.validator.validateLogin(login) ||
            !this.validator.validatePassword(password) ||
            !this.validator.validateName(name) ||
            !this.validator.validateName(surname)
        ) {
            this.setState({
                progress: false,
                showAlert: true,
                password: '',
            });
            return;
        }

        if (await this.userModel.register(name, surname, login, password)) {
            this.setState({ isSuccess: true });
        } else {
            this.setState({
                progress: false,
                showAlert: true,
                password: '',
            });
        }
    };

    render() {
        if (this.state.isSuccess) {
            return <Navigate to={ClientRoutes.profilePage}/>;
        }
        if (!this.props.user) {
            return <Navigate to={ClientRoutes.loginPage} replace={true} />;
        }

        const {
            name,
            surname,
            login,
            password,
            showAlert,
            progress,
        } = this.state;

        return (
            <ContentCard className={registrationCn}>
                <Info
                    show={showAlert}
                    type="alert"
                >
                    Неверный формат ввода одного из полей
                </Info>
                <form onSubmit={this.onSubmit}>
                    <Textinput
                        onChange={this.onChangeName}
                        value={name}
                        label="Имя"
                    />
                    <Textinput
                        onChange={this.onChangeSurname}
                        value={surname}
                        label="Фамилия"
                    />
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
                        className={registrationSubmitCn}
                        view="action"
                        size="m"
                        type="submit"
                        progress={progress}
                    >
                        Зарегистрировать
                    </Button>
                </form>
            </ContentCard>
        );
    }
}
