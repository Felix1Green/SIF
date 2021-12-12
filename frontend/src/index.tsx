import * as React from 'react';
import { render } from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import './index.scss';
import { BaseView } from './views/BaseView';
import { configureRootTheme } from '@yandex/ui/Theme';
import { theme } from './theme/presets/light';

configureRootTheme({ theme });

const App: React.FC = () => {
    return (
        <BrowserRouter >
            <BaseView />
        </BrowserRouter>
    );
};

render(<App />, document.querySelector('.App'));
