import React from 'react';
import { render } from 'react-dom';
import './index.scss';

const App = () => {
    // const hello = 'heloo';

    // console.log(hello);

    return <h1>Hello, world!</h1>;
};

render(<App />, document.getElementById('root'));
