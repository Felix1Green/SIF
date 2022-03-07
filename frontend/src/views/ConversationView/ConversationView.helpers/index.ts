import { Conversation, IStorage, UserStatus, User } from '@chatscope/use-chat';

const defaultAvatar = '/img/avatar-com.svg';

const itkees = {
    avatar: defaultAvatar,
    bio: 'Я банкир',
    firstName: 'Василий',
    lastName: 'Петров',
    username: 'mask',
    email: 'vpetrov@mail.com',
    id: '1',
    presence: {
        status: UserStatus.Available,
        description: 'Оставь мне сообщение',
    },
    role: {
        permissions: [],
    },
};
const putin = {
    avatar: defaultAvatar,
    bio: 'Я президент',
    firstName: 'Владимир',
    lastName: 'Пунин',
    username: 'king',
    email: 'vpunin@mail.com',
    id: '2',
    presence: {
        status: UserStatus.Available,
        description: 'Оставьте меня',
    },
    role: {
        permissions: []
    },
};
export const tintul: User = {
    avatar: defaultAvatar,
    bio: 'Я программист',
    firstName: 'Антон',
    lastName: 'Тинтул',
    username: 'frontend',
    email: 'atintul@mail.com',
    id: '3',
    presence: {
        status: UserStatus.Available,
        description: 'Оставьте меня',
    }
};
const klopp = {
    avatar: defaultAvatar,
    bio: 'I\'m a trainer',
    firstName: 'Юрген',
    lastName: 'Клопп',
    username: 'liverpool',
    email: 'yklopp@mail.com',
    id: '4',
    presence: {
        status: UserStatus.Available,
        description: 'Оставьте меня',
    },
    role: {
        permissions: []
    },
};
const salah = {
    avatar: defaultAvatar,
    bio: 'Я футболист',
    firstName: 'Мохаммед',
    lastName: 'Салах',
    username: 'egypt',
    email: 'msalah@mail.com',
    id: '5',
    presence: {
        status: UserStatus.Available,
        description: 'Оставьте меня',
    },
    role: {
        permissions: []
    },
};
const gorin = {
    avatar: defaultAvatar,
    bio: 'Я фрик',
    firstName: 'Геннадий',
    lastName: 'Горин',
    username: 'goring',
    email: 'goring@mail.com',
    id: '6',
    presence: {
        status: UserStatus.Available,
        description: 'Оставьте меня',
    },
    role: {
        permissions: []
    },
};
const nagiev = {
    avatar: defaultAvatar,
    bio: 'Я ведущий',
    firstName: 'Дмитрий',
    lastName: 'Нагиев',
    username: 'nagiev',
    email: 'dnagiev@mail.com',
    id: '7',
    presence: {
        status: UserStatus.Available,
        description: 'Оставьте меня',
    },
    role: {
        permissions: []
    },
};
const soloveva = {
    avatar: defaultAvatar,
    bio: 'Я репетитор',
    firstName: 'Мария',
    lastName: 'Соловьева',
    username: 'solovevamv',
    email: 'solovevamv@mail.com',
    id: '8',
    presence: {
        status: UserStatus.Available,
        description: 'Оставьте меня',
    },
    role: {
        permissions: []
    },
};

const users = [
    itkees,
    putin,
    klopp,
    salah,
    gorin,
    nagiev,
    soloveva,
];

export const fillUsers = (chatStorage: IStorage) => {
    users.forEach(value => {
        chatStorage.addUser(value);
    });
};

export const fillConversations = (chatStorage: IStorage) => {
    chatStorage.addConversation(new Conversation({
        id: '1',
        participants: [
            { id: itkees.id, role: itkees.role },
            { id: putin.id, role: putin.role },
            { id: salah.id, role: salah.role },
            { id: gorin.id, role: gorin.role },
            { id: nagiev.id, role: nagiev.role },
        ],
    }));
    chatStorage.addConversation(new Conversation({
        id: '2',
        participants: [ { id: putin.id, role: putin.role } ],
    }));
    chatStorage.addConversation(new Conversation({
        id: '3',
        participants: [ { id: salah.id, role: salah.role } ],
    }));
};
