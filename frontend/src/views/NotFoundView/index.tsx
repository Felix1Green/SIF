import * as React from 'react';
import { Link } from 'react-router-dom';
import { notFound404Cn, notFoundCn, notFoundRedirectCn } from './NotFoundView.const';
import { ContentCard } from '@src/components/ContentCard';

import './index.scss';

const NotFoundView: React.FC = () => {
    return (
        <ContentCard className={notFoundCn}>
            <div className={notFound404Cn}>404</div>
            <Link to="/" className={notFoundRedirectCn}>Вернуться на главную</Link>
        </ContentCard>
    );
};

export default NotFoundView;
