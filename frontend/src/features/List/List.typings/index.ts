import { ListItemProps } from '@features/List/List.components/ListItem/ListItem.typings';

export type ListProps = {
    title: string;
    icon: string;
    list?: Array<ListItemProps>;
};
