export type ListItemProps = {
    id?: string;
    url?: string;
    title: string;
    description?: string;
    actions?: ListItemActions;
};

export type ListItemActions = Array<ListItemAction>;

export type ListItemAction = {
    icon: string;
    onClick?: () => void;
}
