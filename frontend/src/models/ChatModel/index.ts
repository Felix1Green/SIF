import {
    ChatMessage,
    Conversation,
    GroupIdGenerator,
    IStorage,
    MessageContentType,
    MessageIdGenerator,
    Participant,
    Presence,
} from '@chatscope/use-chat';
import { ConversationId, GroupedMessages, UserId } from '@chatscope/use-chat/dist/Types';
import { User } from '@chatscope/use-chat/dist/User';
import { MessageGroup } from '@chatscope/use-chat/dist/MessageGroup';

export default class ChatModel implements IStorage {
    private currentUser?: User;
    private users: Array<User>;
    private conversations: Array<Conversation>;
    private messages: GroupedMessages;
    private currentMessage: string;
    private activeConversationId?: ConversationId;
    readonly _groupIdGenerator: GroupIdGenerator;
    readonly _messageIdGenerator?: MessageIdGenerator;

    /**
     * Конструктор
     * @param props
     */
    constructor(props: { groupIdGenerator: GroupIdGenerator, messageIdGenerator: MessageIdGenerator }) {
        this.users = []; // Объект для хранения пулла пользователей чата
        this.conversations = []; // Объект для хранения всех диалогов пользователя
        this.messages = {};
        this.currentMessage = '';
        this._groupIdGenerator = props.groupIdGenerator;
        this._messageIdGenerator = props.messageIdGenerator;
    }
    get groupIdGenerator() {
        return this._groupIdGenerator;
    }
    get messageIdGenerator() {
        return this._messageIdGenerator;
    }

    getMessageWithId(message: ChatMessage<MessageContentType>, generateId: boolean) {
        if (generateId) {
            if (!this.messageIdGenerator) {
                throw 'Id generator is not defined';
            }
            else {
                return Object.assign(Object.assign({}, message), { id: this.messageIdGenerator(message) });
            }
        }
        else {
            return message;
        }
    }

    /**
     * Проверка существования пользователя
     * @param userId
     */
    userExists(userId: UserId) {
        return this.users.findIndex((u) => u.id === userId) !== -1;
    }

    /**
     * Установка залогиненного пользователя
     * @param user
     */
    setCurrentUser(user: User) {
        this.currentUser = user;
    }

    /**
     * Добавление пользователя в колекцию пользователей
     * @param user
     */
    addUser(user: User) {
        const notExists = !this.userExists(user.id);
        if (notExists) {
            this.users = this.users.concat(user);
        }
        return notExists;
    }
    /**
     * Удаление пользователя из коллекции пользователей
     * @param userId
     */
    removeUser(userId: UserId) {
        const idx = this.users.findIndex((u) => u.id === userId);
        if (idx !== -1) {
            this.users = this.users.slice(0, idx).concat(this.users.slice(idx + 1));
            return true;
        }
        return false;
    }
    /**
     * Получение пользователя по его ID
     * @param userId
     * @return [User, number]|[undefined,undefined]
     */
    getUser(userId: UserId) : [User, number] | [undefined, undefined] {
        const idx = this.users.findIndex((u) => u.id === userId);
        if (idx !== -1) {
            return [ this.users[idx], idx ];
        }
        return [ undefined, undefined ];
    }
    /**
     * Проверка существования диалога
     * @param conversationId
     */
    conversationExists(conversationId: ConversationId) {
        return this.conversations.findIndex((c) => c.id === conversationId) !== -1;
    }
    /**
     * Получение диалога по его ID
     * @param conversationId
     * @return [Conversation, number]|[undefined, undefined]
     */
    getConversation(conversationId: ConversationId): [Conversation, number] | [undefined, undefined]{
        const idx = this.conversations.findIndex((c) => c.id === conversationId);
        if (idx !== -1) {
            return [ this.conversations[idx], idx ];
        }
        return [ undefined, undefined ];
    }
    /**
     * Добавление диалога в коллекцию чатов
     * @param conversation
     */
    addConversation(conversation: Conversation) {
        const notExists = !this.conversationExists(conversation.id);
        if (notExists) {
            this.conversations = this.conversations.concat(conversation);
        }
        return notExists;
    }
    /**
     * Установка количества непрочитанных сообщений
     * @param conversationId
     * @param count
     */
    setUnread(conversationId: ConversationId, count: number) {
        const [ conversation ] = this.getConversation(conversationId);
        if (conversation) {
            conversation.unreadCounter = count;
        }
    }

    /**
     * Удаление диалога из коллекции диалогов
     * @param conversationId
     * @param removeMessages
     */
    removeConversation(conversationId: ConversationId, removeMessages = true) {
        const idx = this.conversations.findIndex((c) => c.id === conversationId);
        if (idx !== -1) {
            this.conversations = this.conversations
                .slice(0, idx)
                .concat(this.conversations.slice(idx + 1));
            if (removeMessages) {
                delete this.messages[conversationId];
            }
            return true;
        }
        return false;
    }

    /**
     * Замена диалога другим диалогом
     * @param conversation
     * @param idx
     */
    replaceConversation(conversation: Conversation, idx: number) {
        this.conversations = this.conversations
            .slice(0, idx)
            .concat(new Conversation({
                id: conversation.id,
                participants: conversation.participants,
                typingUsers: conversation.typingUsers,
                unreadCounter: conversation.unreadCounter,
                draft: conversation.draft,
                description: conversation.description,
                readonly: conversation.readonly,
            }))
            .concat(this.conversations.slice(idx + 1));
    }

    /**
     * Замена пользователя другим пользователем
     * @param user
     * @param idx
     */
    replaceUser(user: User, idx: number) {
        this.users = this.users
            .slice(0, idx)
            .concat(user)
            .concat(this.users.slice(idx + 1));
    }

    /**
     * Добавление участника в диалог
     * @param conversationId
     * @param participant
     * @return boolean
     */
    addParticipant(conversationId: ConversationId, participant: Participant) {
        const [ conversation, idx ] = this.getConversation(conversationId);
        if (conversation && idx) {
            if (conversation.addParticipant(participant)) {
                this.replaceConversation(conversation, idx);
            }
        }
        return false;
    }
    /**
     * Удаление участника из диалога
     * @param conversationId
     * @param participantId
     */
    removeParticipant(conversationId: ConversationId, participantId: string) {
        const [ conversation, idx ] = this.getConversation(conversationId);
        if (conversation && idx) {
            conversation.removeParticipant(participantId);
            this.replaceConversation(conversation, idx);
            return true;
        }
        return false;
    }

    /**
     * Отправить сообщение
     * @param message
     * @param conversationId
     * @param generateId
     */
    addMessage(message: ChatMessage<MessageContentType>, conversationId: ConversationId, generateId = false) {
        if (conversationId in this.messages) { // Поиск id диалога в объекте диалогов и их сообщений
            const groups = this.messages[conversationId]; // Выделяем массив сообщений конркетного диалога
            const lastGroup = groups[groups.length - 1]; // Поиск id диалога в объекте диалогов и их сообщений
            if (lastGroup.senderId === message.senderId) {
                // Add message to group
                const newMessage = this.getMessageWithId(message, generateId);
                lastGroup.addMessage(newMessage);
                return newMessage;
            }
        }
        const group = new MessageGroup({
            id: this.groupIdGenerator(),
            senderId: message.senderId,
            direction: message.direction,
        });
        const newMessage = this.getMessageWithId(message, generateId);
        group.addMessage(newMessage);
        this.messages[conversationId] =
            conversationId in this.messages
                ? this.messages[conversationId].concat(group)
                : [ group ];
        return newMessage;
    }

    // TODO: Refactoring - it's not very optimal :)
    /**
     * Обновить сообщение
     * @param message
     */
    updateMessage(message: ChatMessage<MessageContentType>) {
        for (const conversationId in this.messages) {
            if (!Object.prototype.hasOwnProperty.call(this.messages, conversationId)) {
                continue;
            }
            const groups = this.messages[conversationId];
            const l = groups.length;
            for (let i = 0; i < l; i++) {
                const group = groups[i];
                const [ currentMessage, idx ] = group.getMessage(message.id);
                if (currentMessage && idx) {
                    group.replaceMessage(message, idx);
                }
            }
        }
    }

    /**
     * Установить тип доступности пользователя
     * @param presence
     */
    setPresence(userId: UserId, presence: Presence) {
        const [ user, idx ] = this.getUser(userId);
        if (user && idx) {
            user.presence = presence;
            this.replaceUser(user, idx);
        }
    }

    /**
     * Установить черновик сообщения для текущего диалога
     * @param {String} draft
     */
    setDraft(draft: string) {
        if (this.activeConversationId) {
            const [ activeConversation, idx ] = this.getConversation(this.activeConversationId);
            if (activeConversation && idx) {
                activeConversation.draft = draft;
                this.replaceConversation(activeConversation, idx);
            }
        }
    }

    /**
     * Получение состояния чата
     */
    getState() {
        return {
            currentUser: this.currentUser,
            users: this.users,
            conversations: this.conversations,
            // TODO: Implement sth like collection referencing by id (map with guaranteed order like in immutablejs) because searching every time in the array does not make sense
            activeConversation: this.activeConversationId
                ? this.conversations.find((c) => c.id === this.activeConversationId)
                : undefined,
            currentMessages: this.activeConversationId && this.activeConversationId in this.messages
                ? this.messages[this.activeConversationId]
                : [],
            messages: this.messages,
            currentMessage: this.currentMessage,
        };
    }

    /**
     * Сбрасывает состояние чата
     */
    clearState() { }
    resetState() {
        this.currentUser = undefined;
        this.users = [];
        this.conversations = [];
        this.activeConversationId = undefined;
        this.messages = {};
    }
    /**
     * Устанавливает текущий активный диалог и сбрасывает счетчик непрочитанных сообщений диалога, если второй параметр установлен
     * @param conversationId
     * @param resetUnreadCounter
     */
    setActiveConversation(conversationId?: ConversationId, resetUnreadCounter = true) {
        this.activeConversationId = conversationId;
        if (resetUnreadCounter && conversationId) {
            const [ conversation, idx ] = this.getConversation(conversationId);
            if (conversation && idx) {
                conversation.unreadCounter = 0;
                this.replaceConversation(conversation, idx);
            }
        }
    }
    /**
     * Устанавливает текущее сообщение из инпута
     * @param message
     */
    setCurrentMessage(message: string) {
        this.currentMessage = message;
    }
}
