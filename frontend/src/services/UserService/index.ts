class UserService {
    async login(login?: string, password?: string) {
        return new Promise<true>((resolve, reject) => {
            if (!login || !password) {
                return reject(false);
            }
            resolve(true);
        });
    }
}

export default UserService;
