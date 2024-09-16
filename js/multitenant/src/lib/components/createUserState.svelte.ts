export default function createUserState(initial: { username: string, password: string }) {
  let username = $state(initial.username);
  let password = $state(initial.password);
  return {
    ...initial,
    get password() {
      return password;
    },
    get username() {
      return username;
    },
    set password(value) {
      password = value;
    },
    set username(value) {
      username = value;
    },
  };
}
