import React, { useContext, useEffect } from "react";
import Navbar from "../Navbar";
import { Outlet, useNavigate } from "react-router";
import { AuthContext } from "../../context/auth";

const HomeLayout: React.FC = () => {
  const { auth } = useContext(AuthContext) || { auth: null };
  const navigate = useNavigate();

  useEffect(() => {
    if (!auth?.email) {
      navigate("/login");
    }
  }, [auth]);
  return (
    <main className="w-full h-screen flex flex-col bg-slate-200">
      <Navbar></Navbar>
      <div className="w-full h-full ">
        <Outlet />
      </div>
    </main>
  );
};

export default HomeLayout;
