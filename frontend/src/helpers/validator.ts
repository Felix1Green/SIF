export default class Validator {
    validateLogin(login?: string): boolean {
        if (!login) {
            return false;
        }
        const loginRegExp = /^[a-z0-9._%+-]+@[a-z0-9-]+.+.[a-z]{2,4}$/;
        return loginRegExp.test(login);
    }

    validatePassword(password?: string): boolean {
        if (!password) {
            return false;
        }
        const passwordRegExp = /^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[@#$%^&+=])(?=\S+$).{8,}$/;
        return passwordRegExp.test(password);
    }

    validateName(name?: string): boolean {
        if (!name) {
            return false;
        }
        const nameRegExp = /^([A-Z][a-z]*)|([А-Я][а-я]*)$/;
        return nameRegExp.test(name);
    }

    validateSurname(surname?: string): boolean {
        return this.validateName(surname);
    }

    validateRole(role?: string): boolean {
        return Boolean(role);
    }
}
