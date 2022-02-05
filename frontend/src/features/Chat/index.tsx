import * as React from 'react';
import {
    useState,
    useCallback
} from 'react';
import {
    MainContainer,
    Sidebar,
    ConversationList,
    Conversation,
    MessageGroup,
    Message,
    ChatContainer,
    ConversationHeader,
    MessageList,
    MessageInput,
    Avatar,
    AvatarGroup,
} from '@chatscope/chat-ui-kit-react';
import {
    useChat,
    ChatMessage,
    MessageContentType,
    MessageDirection,
    MessageStatus, Conversation as ConversationType,
} from '@chatscope/use-chat';
import { User } from '@chatscope/use-chat/dist/User';

type ChatState = {
    acAvatar: JSX.Element | Array<JSX.Element> | undefined;
    acUsername: string | undefined;
};

export type ChatProps = {
    user: User;
};

export const Chat = (props: ChatProps) => {
    const { user } = props;
    const [ messageInput, setMessageInput ] = useState('');

    // Get all chat related values and methods from useChat hook
    const {
        currentMessages,
        conversations,
        activeConversation,
        setActiveConversation,
        sendMessage,
        getUser
    } = useChat();

    // ac - active conversation
    const [ acHeader, setACHeaderState ] = useState<ChatState>({
        acAvatar: undefined,
        acUsername: undefined,
    });
    const setACHeaderData = useCallback((conversationId: string) => {
        setActiveConversation(conversationId);
        const ac = conversations.filter(value => value.id === conversationId)[0];
        if (!ac) {
            return null;
        }

        const { avatar, username } = getConversationData(ac);
        setACHeaderState({ acAvatar: avatar, acUsername: username });
    }, [ activeConversation ]);

    const handleMessageSend = (text: any) => {
        if (!activeConversation) return;

        // Logger user (sender)
        const senderId = user.id;
        const message = new ChatMessage({
            id: '',
            content: text,
            contentType: MessageContentType.TextHtml,
            senderId,
            direction: MessageDirection.Outgoing,
            status: MessageStatus.Sent
        });
        sendMessage({
            message,
            conversationId: activeConversation.id,
            senderId,
        });
        setMessageInput('');
    };

    const getConversationData = (conversation: ConversationType) => {
        let avatar, username;
        if (conversation.participants.length > 1) {
            avatar = (
                <AvatarGroup size="sm" style={{ width: '48px' }}>
                    {conversation.participants.map((value, index) => {
                        const user = getUser(value.id);
                        if (user && index < 4) {
                            return <Avatar key={index} src={user.avatar} name={user.firstName}/>;
                        }
                    })}
                </AvatarGroup>
            );
            username = conversation.participants.map(value => {
                const user = getUser(value.id);
                if (user) {
                    return `${user.firstName} ${user.lastName}`;
                }
            }).join(', ');
        } else {
            const user = getUser(conversation.participants[0].id);
            if (user) {
                avatar = <Avatar src={user.avatar}/>;
                username = `${user.firstName} ${user.lastName}`;
            }
        }
        return { avatar, username };
   };
    const getConversationsList = () => {
        return conversations.map(conversation => {
            const { avatar, username } = getConversationData(conversation);

            return (
                <Conversation
                    key={conversation.id}
                    name={username}
                    active={activeConversation?.id === conversation.id}
                    unreadCnt={conversation.unreadCounter}
                    onClick={() => setACHeaderData(conversation.id)}
                >
                    {avatar}
                </Conversation>
            );
        });
    };

    return (
        <MainContainer>
            <Sidebar position="left">
                <ConversationHeader>
                    <Avatar src={user.avatar} />
                    <ConversationHeader.Content>
                        {`${user.firstName} ${user.lastName}`}
                    </ConversationHeader.Content>
                </ConversationHeader>
                <ConversationList>
                    {getConversationsList()}
                </ConversationList>
            </Sidebar>
            <ChatContainer>
                <ConversationHeader>
                    {acHeader.acAvatar}
                    <ConversationHeader.Content
                        userName={acHeader.acUsername}
                    />
                </ConversationHeader>
                <MessageList>
                    {currentMessages.map(g => (
                        <MessageGroup key={g.id} direction={g.direction}>
                            <MessageGroup.Messages>
                                {g.messages.map(m => (
                                    <Message
                                        key={m.id}
                                        model={{
                                            type: 'text',
                                            payload: m.content
                                        }}
                                    />
                                ))}
                            </MessageGroup.Messages>
                        </MessageGroup>
                    ))}
                </MessageList>
                <MessageInput
                    placeholder="Сообщение"
                    value={messageInput}
                    onSend={handleMessageSend}
                    onChange={setMessageInput}
                    attachButton={false}
                />
            </ChatContainer>
        </MainContainer>
    );
};
