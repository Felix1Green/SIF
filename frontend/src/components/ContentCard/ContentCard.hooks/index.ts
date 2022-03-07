import { useState } from 'react';

export const useCollapse = (collapsed: boolean | undefined): [ boolean | undefined, (() => void) | undefined] => {
    if (collapsed === undefined) {
        return [ undefined, undefined ];
    }
    const [ collapseState, setCollapseState ] = useState(collapsed);
    const onCollapse = () => { setCollapseState(state => !state); };

    return [ collapseState, onCollapse ];
};
