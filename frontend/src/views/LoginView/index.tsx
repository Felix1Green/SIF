import * as React from 'react';
import { ChangeEvent, FormEvent, useState } from 'react';
import { LoginViewPropsType } from './LoginView.typings';
import { inputCn, labelCn, loginCn, submitCn } from './LoginView.consts';
import { Textinput } from '@yandex/ui/Textinput/desktop/bundle';
import { Button } from '@yandex/ui/Button/desktop/bundle';
import { Text } from '@yandex/ui/Text/desktop/bundle';
import { ContentCard } from '@components/ContentCard';
import UserService from '@services/UserService';

import './index.scss';
import Validator from '@helpers/validator';
import { Info } from '@components/Info';

const LoginView: React.FC<LoginViewPropsType> = () => {
    const userService = new UserService();
    const validator = new Validator();
    const [ login, setLogin ] = useState('');
    const [ password, setPassword ] = useState('');
    const [ showAlert, setShowAlert ] = useState<boolean>(false);
    const [ progress, setProgress ] = useState<boolean>(false);

    const onSubmit = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setProgress(true);
        setShowAlert(false);
        await new Promise((resolve) => setTimeout(async () => {resolve('hello');}, 2000));
        if (!validator.validateLogin(login) || !validator.validatePassword(password)) {
            setProgress(false);
            setShowAlert(true);
            setPassword('');
            return;
        }
        userService.logIn(login, password);
    };

    const onChangeLogin = (event: ChangeEvent<HTMLInputElement>) => {
        setLogin(event.target.value);
    };
    const onChangePassword = (event: ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
    };

    return (
        <ContentCard className={loginCn}>
            <Info
                show={showAlert}
                type="alert"
            >
                Неверный логин и/или пароль
            </Info>
            <form onSubmit={onSubmit}>
                <Text className={labelCn}>Логин</Text>
                <Textinput
                    onChange={onChangeLogin}
                    value={login}
                    className={inputCn}
                    size="m"
                    view="default"
                    required
                />
                <Text className={labelCn}>Пароль</Text>
                <Textinput
                    label="Пароль"
                    onChange={onChangePassword}
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
};

export default LoginView;
