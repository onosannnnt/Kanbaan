import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { axiosInstance } from "../utils/axios";
import { useContext, useEffect } from "react";
import { useNavigate } from "react-router";
import { AuthContext } from "../context/auth";

const schema = z.object({
  email: z.string().email(),
  password: z
    .string()
    .min(8, { message: "Password must be at least 8 characters long" }),
});

type FormData = z.infer<typeof schema>;

const Login = () => {
    const navigate = useNavigate();
    const { auth } = useContext(AuthContext) || { auth: null };
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(schema),
  });

    useEffect(() => {
      if (auth?.email) {
        navigate("/");
      }
    }, [auth]);

  const onSubmit = async(data: FormData) => {
    try{
        const response = await axiosInstance.post("/users/login", data);
        if (response.status === 200) {
            console.log("Login successful:", response.data);
            // Handle successful login, e.g., redirect or show a success message
            window.location.href = "/";
        } else {
            console.error("Login failed:", response.data);
            // Handle login failure, e.g., show an error message
        }
    }catch (error) {
        console.error("Error during login:", error);
    }
  };

  return (
    <main className="flex flex-col items-center justify-center w-full h-screen">
      <div className="w-1/4 p-8 bg-white shadow-lg rounded flex flex-col items-center gap-4">
        <h1 className="text-2xl font-bold mb-4">Kanban Clicknext</h1>
        <h2 className="text-2xl font-bold mb-4">Login</h2>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="flex flex-col gap-8 w-full"
        >
          <div className="flex flex-col">
            <label className="text-m font-bold mb-2">Email</label>
            <input
              {...register("email")}
              placeholder="Email"
              className="p-2 w-full border-b-2 border-gray-300 rounded"
            />
            {errors.email && (
              <p className="text-red-500">{errors.email.message}</p>
            )}
          </div>
          <div className="flex flex-col">
            <label className="text-m font-bold mb-2">Password</label>
            <input
              {...register("password")}
              type="password"
              placeholder="Password"
              className="p-2 w-full border-b-2 border-gray-300 rounded"
            />
            {errors.password && (
              <p className="text-red-500">{errors.password.message}</p>
            )}
          </div>
          <button
            type="submit"
            className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600 transition duration-200"
          >
            Login
          </button>
          <div>
            <p className="text-sm">
              Don't have an account?{" "}
              <a href="/register" className="text-blue-500 hover:text-blue-700">
                Register
              </a>
            </p>
          </div>
        </form>
      </div>
    </main>
  );
};
export default Login;
