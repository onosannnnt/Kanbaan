import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";

const schema = z
  .object({
    email: z.string().email(),
    password: z
      .string()
      .min(8, { message: "Password must be at least 8 characters long" }),
    comfirmPassword: z
      .string()
      .min(8, { message: "Password must be at least 8 characters long" }),
    firstname: z.string().min(1, { message: "First name cannot be null" }),
    lastname: z.string().min(1, { message: "Last name cannot be null" }),
  })
  .superRefine((data, ctx) => {
    if (data.comfirmPassword !== data.password) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: "Passwords do not match",
        path: ["comfirmPassword"],
      });
    }
  });

type FormData = z.infer<typeof schema>;

const Register = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(schema),
  });

  const onSubmit = (data: FormData) => {
    console.log("Submitted:", data);
  };

  return (
    <main className="flex flex-col items-center justify-center w-full h-screen">
      <div className="w-1/4 p-8 bg-white shadow-lg rounded flex flex-col items-center gap-4">
        <h1 className="text-2xl font-bold mb-4">Kanban Clicknext</h1>
        <h2 className="text-2xl font-bold mb-4">Register</h2>
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
            <label className="text-m font-bold mb-2">First Name</label>
            <input
              {...register("firstname")}
              placeholder="First Name"
              className="p-2 w-full border-b-2 border-gray-300 rounded"
            />
            {errors.firstname && (
              <p className="text-red-500">{errors.firstname.message}</p>
            )}
          </div>
          <div className="flex flex-col">
            <label className="text-m font-bold mb-2">Last Name</label>
            <input
              {...register("lastname")}
              placeholder="Last Name"
              className="p-2 w-full border-b-2 border-gray-300 rounded"
            />
            {errors.lastname && (
              <p className="text-red-500">{errors.lastname.message}</p>
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
          <div className="flex flex-col">
            <label className="text-m font-bold mb-2">Confirm Password</label>
            <input
              {...register("comfirmPassword")}
              type="password"
              placeholder="Confirm Password"
              className="p-2 w-full border-b-2 border-gray-300 rounded"
            />
            {errors.comfirmPassword && (
              <p className="text-red-500">{errors.comfirmPassword.message}</p>
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
              Already have an account?{" "}
              <a href="/login" className="text-blue-500 hover:text-blue-700">
                Login
              </a>
            </p>
          </div>
        </form>
      </div>
    </main>
  );
};
export default Register;
