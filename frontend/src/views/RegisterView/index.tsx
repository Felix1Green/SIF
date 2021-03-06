import * as React from 'react';
import { ChangeEvent, FormEvent } from 'react';
import { RegisterViewProps, RegisterViewState } from '@views/RegisterView/RegisterView.typings';
import Validator from '@helpers/validator';
import UserModel from '@models/UserModel';
import { ClientRoutes } from '@consts/routes';
import { Navigate } from 'react-router-dom';
import { ContentCard } from '@src/components/ContentCard';
import { registrationCn, registrationFormCn, registrationSubmitCn, roleSelectOptions } from './RegisterView.const';
import { Roles } from '@views/ProfileView/ProfileView.typings';
import { Disclaimer } from '@components/Disclaimer';
import { Textinput } from '@components/Textinput';
import { Button } from '@yandex/ui/Button/desktop/bundle';
import { Select } from '@components/Select';

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
            role: Roles.Default,
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

    onChangeRole = (event: ChangeEvent<HTMLSelectElement>) => {
        this.setState({ role: event.target.value as Roles });
    };

    onSubmit = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const {
            login,
            password,
            name,
            surname,
            role,
        } = this.state;

        this.setState({
            progress: true,
            showAlert: false,
        });

        // TODO: ???????????????????? ???????????????????????? ???????????? ?????? ???????????????????????? ????????????????
        // await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 1000));

        if (
            !this.validator.validateLogin(login) ||
            !this.validator.validatePassword(password) ||
            !this.validator.validateName(name) ||
            !this.validator.validateSurname(surname) ||
            !this.validator.validateRole(role)
        ) {
            this.setState({
                progress: false,
                showAlert: true,
                password: '',
            });
            return;
        }

        if (await this.userModel.register({ name, surname, role, login, password })) {
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
            role,
            login,
            password,
            showAlert,
            progress,
        } = this.state;

        return (
            <ContentCard className={registrationCn} title="??????????????????????">
                <Disclaimer
                    show={showAlert}
                    type="alert"
                >
                    ???????????????? ???????????? ?????????? ???????????? ???? ??????????
                </Disclaimer>
                <form className={registrationFormCn} onSubmit={this.onSubmit}>
                    <Textinput
                        onChange={this.onChangeName}
                        value={name}
                        label="??????"
                    />
                    <Textinput
                        onChange={this.onChangeSurname}
                        value={surname}
                        label="??????????????"
                    />
                    <Textinput
                        onChange={this.onChangeLogin}
                        value={login}
                        label="??????????"
                    />
                    <Select
                        size="m"
                        view="default"
                        label="????????"
                        onChange={this.onChangeRole}
                        value={role}
                        placeholder={'?????????????? ????????'}
                        options={roleSelectOptions}
                    />
                    <Textinput
                        onChange={this.onChangePassword}
                        value={password}
                        label="????????????"
                        type="password"
                    />
                    <Button
                        className={registrationSubmitCn}
                        view="action"
                        size="m"
                        type="submit"
                        progress={progress}
                    >
                        ????????????????????????????????
                    </Button>
                </form>
            </ContentCard>
        );
    }
}
