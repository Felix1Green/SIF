declare module '@chatscope/chat-ui-kit-react' {
    import * as React from 'react';

    type Avatar = React.ComponentClass<any>
    export const Avatar: Avatar;

    type AvatarGroup = React.ComponentClass<any>
    export const AvatarGroup: AvatarGroup;

    type ChatContainer = React.ComponentClass<any>
    export const ChatContainer: ChatContainer;

    interface ConversationHeader extends React.ComponentClass<any> {
        Back: any
        Content: any
        Actions: any
    }
    export const ConversationHeader: ConversationHeader;

    type MainContainer = React.ComponentClass<any>
    export const MainContainer: MainContainer;

    type Sidebar = React.ComponentClass<any>
    export const Sidebar: Sidebar;

    type ConversationList = React.ComponentClass<any>
    export const ConversationList: ConversationList;

    type Conversation = React.ComponentClass<any>
    export const Conversation: Conversation;

    interface MessageGroup extends React.ComponentClass<any> {
        Messages: any
    }
    export const MessageGroup: MessageGroup;

    type Message = React.ComponentClass<any>
    export const Message: Message;

    type MessageList = React.ComponentClass<any>
    export const MessageList: MessageList;

    type MessageInput = React.ComponentClass<any>
    export const MessageInput: MessageInput;
}
