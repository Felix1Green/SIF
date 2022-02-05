import * as React from 'react';
import { render } from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import './index.scss';
import { BaseView } from '@views/BaseView';
import { configureRootTheme } from '@yandex/ui/Theme';
import { theme } from '@theme/presets/light';
import { UserWrapper } from '@views/ProfileView/ProfileView.typings';
import UserModel from '@models/UserModel';
import { Spin } from '@yandex/ui/Spin/desktop/bundle';

configureRootTheme({ theme });

export type AppState = {
    user: UserWrapper,
    setUser: (recievedUser: UserWrapper) => void,
};

export type AppProps = Record<string, never>;

export const UserContext = React.createContext<AppState>({
    user: undefined,
    setUser: () => {},
});

class App extends React.Component<AppProps, AppState> {
    public setUser: (receivedUser: UserWrapper) => void;
    private userModel: UserModel;

    constructor(props: AppProps) {
        super(props);
        this.userModel = new UserModel();

        this.setUser = (receivedUser: UserWrapper) => {
            if (receivedUser) {
                this.setState({ user: receivedUser });
            } else {
                this.setState({ user: null });
            }
        };

        this.state = {
            user: undefined,
            setUser: this.setUser,
        };
    }

    render() {
        if (this.state.user === undefined) {
            this.userModel.getUserProfile().then(result => {
                this.setUser(result);
            });

            return <Spin progress view="default" size="l" />;
        }

        return (
            <BrowserRouter>
                <UserContext.Provider value={this.state}>
                    <BaseView />
                </UserContext.Provider>
            </BrowserRouter>
        );
    }
}

render(<App />, document.querySelector('.App'));
