import * as React from 'react';
import { Navigate } from 'react-router-dom';
import '@chatscope/chat-ui-kit-styles/dist/default/styles.min.css';
import { nanoid } from 'nanoid';
import {
    IStorage,
    AutoDraft,
    UpdateState,
    ChatProvider,
    ChatServiceFactory,
    GroupIdGenerator,
    MessageIdGenerator,
} from '@chatscope/use-chat';
import { ExampleChatService } from '@chatscope/use-chat/dist/examples';
import { IChatService } from '@chatscope/use-chat/dist/interfaces/IChatService';
import { ConversationViewProps, ConversationViewState } from '@views/ConversationView/ConversationView.typings';
import { ContentCard } from '@src/components/ContentCard';
import { Chat } from '@features/Chat';
import { ClientRoutes } from '@consts/routes';
import ChatModel from '@models/ChatModel';

import { fillConversations, fillUsers, tintul } from './ConversationView.helpers';
import { conversationWrapperCn } from './ConversationView.const';

import './index.scss';

export default class ConversationView extends React.Component<ConversationViewProps, ConversationViewState> {
    serviceFactory: ChatServiceFactory<IChatService>;
    chatStorage: ChatModel;
    messageIdGenerator: MessageIdGenerator;
    groupIdGenerator: GroupIdGenerator;

    constructor(props: ConversationViewProps) {
        super(props);

        this.serviceFactory = (storage: IStorage, updateState: UpdateState) => {
            return new ExampleChatService(storage, updateState);
        };

        this.messageIdGenerator = () => nanoid();
        this.groupIdGenerator = () => nanoid();
        this.chatStorage = new ChatModel({
            groupIdGenerator: this.groupIdGenerator,
            messageIdGenerator: this.messageIdGenerator,
        });

        fillUsers(this.chatStorage);
        fillConversations(this.chatStorage);
    }

    render() {
        if (!this.props.user) {
            return <Navigate to={ClientRoutes.loginPage} replace={true}/>;
        }

        return (
            <ContentCard className={conversationWrapperCn}>
                <ChatProvider
                    serviceFactory={this.serviceFactory}
                    storage={this.chatStorage}
                    config={{
                        typingThrottleTime: 250,
                        typingDebounceTime: 900,
                        debounceTyping: true,
                        autoDraft: AutoDraft.Save | AutoDraft.Restore
                    }}
                >
                    <Chat user={tintul}/>
                </ChatProvider>
            </ContentCard>
        );
    }
}
