import { useState } from "react";

import { useForm, SubmitHandler } from "react-hook-form";
import { Login, Register } from "../actions/user/register";
import { useAppDispatch } from "../hooks/hook";
type LoginPayload = {
  email: string;
  password: string;
};

type RegisterPayload = {
  email: string;
  username: string;
  password: string;
};

const LoginRegister = () => {
  const dispatch = useAppDispatch();

  const loginForm = useForm<LoginPayload>();
  const registerForm = useForm<RegisterPayload>();

  const [isLogin, setIsLogin] = useState(true);

  const loginHandler: SubmitHandler<LoginPayload> = (data) => {
    console.log(data);
    try {
      console.log(data);
      const payload: LoginPayload = {
        email: data.email,
        password: data.password,
      };
      dispatch(Login(payload));
    } catch (error) {
      console.log(error);
    }
  };

  const registerHandler: SubmitHandler<RegisterPayload> = async (data) => {
    try {
      console.log(data);
      const payload: RegisterPayload = {
        username: data.username,
        email: data.email,
        password: data.password,
      };
      const d = dispatch(Register(payload));
      console.log(d.unwrap());
    } catch (error) {
      console.log(error);
    }
  };

  const toggleForm = () => {
    setIsLogin(!isLogin);
  };

  return (
    <div
      className="bg-cover bg-center bg-fixed"
      style={{ backgroundImage: `url('https://picsum.photos/1920/1080')` }}
    >
      <div className="h-screen flex justify-center items-center">
        <div className="bg-white mx-4 p-8 rounded shadow-md w-full md:w-1/2 lg:w-1/3">
          <h1 className="text-3xl font-bold mb-8 text-center">
            {isLogin ? "Login" : "Register"}
          </h1>
          <form
            onSubmit={
              isLogin
                ? loginForm.handleSubmit(loginHandler)
                : registerForm.handleSubmit(registerHandler)
            }
          >
            {isLogin ? (
              <>
                <div className="mb-4">
                  <label
                    className="block font-semibold text-gray-700 mb-2"
                    htmlFor="email"
                  >
                    Email Address
                  </label>
                  <input
                    className="border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    id="email"
                    type="email"
                    placeholder="Enter your email address"
                    {...loginForm.register("email")}
                  />
                </div>
                <div className="mb-4">
                  <label
                    className="block font-semibold text-gray-700 mb-2"
                    htmlFor="password"
                  >
                    Password
                  </label>
                  <input
                    className="border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
                    id="password"
                    type="password"
                    placeholder="Enter your password"
                    {...loginForm.register("password")}
                  />
                  <a className="text-gray-600 hover:text-gray-800" href="#">
                    Forgot your password?
                  </a>
                </div>
              </>
            ) : (
              <>
                <div className="mb-4">
                  <label
                    className="block font-semibold text-gray-700 mb-2"
                    htmlFor="username"
                  >
                    Username
                  </label>
                  <input
                    className="border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    id="username"
                    type="text"
                    placeholder="Enter your username"
                    {...registerForm.register("username")}
                  />
                </div>
                <div className="mb-4">
                  <label
                    className="block font-semibold text-gray-700 mb-2"
                    htmlFor="email"
                  >
                    Email Address
                  </label>
                  <input
                    className="border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    id="email"
                    type="email"
                    {...registerForm.register("email")}
                    placeholder="Enter your email address"
                  />
                </div>
                <div className="mb-4">
                  <label
                    className="block font-semibold text-gray-700 mb-2"
                    htmlFor="password"
                  >
                    Password
                  </label>
                  <input
                    className="border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
                    id="password"
                    type="password"
                    placeholder="Enter your password"
                    {...registerForm.register("password")}
                  />
                </div>
              </>
            )}
            <div className="mb-6">
              <button
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                type="submit"
              >
                {isLogin ? "Login" : "Register"}
              </button>
            </div>
          </form>
          <div className="text-center">
            <p>
              {isLogin ? "Don't have an account?" : "Already have an account?"}
              <span
                className="text-blue-500 cursor-pointer"
                onClick={toggleForm}
              >
                {isLogin ? " Register" : " Login"}
              </span>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginRegister;
