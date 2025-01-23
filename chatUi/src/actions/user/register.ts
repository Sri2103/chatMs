import Api from "../../utils/api";

export const userApi = new Api("http://localhost:7080/user")

export type RegisterPayload = {
	username: string;
	email: string;
	password: string;
}
export async function Register(data: RegisterPayload) {
	const resp = await userApi.post("/register", JSON.stringify(data))
	return resp
}

export type LoginPayload = {
	email: string;
	password: string;
}
export async function Login(data: LoginPayload) {
	const resp = await userApi.post("/login", JSON.stringify(data))
	return resp
}



