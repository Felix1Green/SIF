import * as React from 'react';
import { render } from 'react-dom';
import './index.scss';
import { BaseView } from './views/BaseView';

const App: React.FC = () => {
    return <BaseView />;
};

render(<App />, document.getElementById('root'));
