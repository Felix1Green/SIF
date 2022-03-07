import * as React from 'react';

import { useMemo } from 'react';

import { classnames } from '@bem-react/classnames';

import { TitleProps } from '@components/Title/Title.typings';
import { titleCn } from '@components/Title/Title.const';

import './index.scss';

export const Title: React.FC<TitleProps> = (props) => {
    const { children, className } = props;
    const titleClassName = useMemo(() => className ? classnames(titleCn, className): titleCn, [ className ]);

    return (
        <div className={titleClassName}>{children}</div>
    );
};
