import { BrowserRouter, Route, Routes } from "react-router";
import HomeLayout from "./components/Layout/HomeLayout";
import Homepage from "./pages/Homepage";
import { AuthContext, InitAuthValue, type IContextType } from "./context/auth";
import { useCallback, useEffect, useState } from "react";
import { axiosInstance } from "./utils/axios";
import Loading from "./components/Loading";

import Register from "./pages/Register";
import Board from "./components/Board";
import Login from "./pages/Login";

const App = () => {
  const [auth, setAuth] = useState<IContextType>(InitAuthValue);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const handleLogin = useCallback(async () => {
    try {
      setIsLoading(true);
      const response = await axiosInstance.get("/users/me");
      setAuth({
        id: response.data.data.ID,
        email: response.data.data.Email,
        firstName: response.data.data.FirstName,
        lastName: response.data.data.LastName,
      });
    } catch {
      setAuth(InitAuthValue);
    } finally {
      setIsLoading(false);
    }
  }, []);

  useEffect(() => {
    handleLogin().then(() => console.log("success"));
  }, [handleLogin]);

  useEffect(() => {
    console.log("auth", auth);
  }, [auth]);

  if (isLoading) {
    return (
      <>
        <div className="w-full h-screen content-center text-center">
          <Loading />
        </div>
      </>
    );
  }
  return (
    <AuthContext.Provider value={{ auth, setAuth }}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomeLayout />}>
            <Route index element={<Homepage />} />
            <Route path="/board/:id" element={<Board />} />
            <Route path="/notification" element={<Loading></Loading>} />
          </Route>
          <Route path="/login" element={<Login></Login>} />
          <Route path="/register" element={<Register></Register>} />
        </Routes>
      </BrowserRouter>
    </AuthContext.Provider>
  );
};

export default App;
