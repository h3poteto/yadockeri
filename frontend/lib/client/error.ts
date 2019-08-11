export class AuthenticationError extends Error {
  constructor(msg: string) {
    super(msg)

    Object.setPrototypeOf(this, AuthenticationError.prototype)
  }
}
