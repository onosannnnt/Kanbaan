import { createContext } from "react";

export interface IContextType {
  id: string;
  email: string;
  firstName?: string;
  lastName?: string;
}

export const InitAuthValue: IContextType = {
  id: "",
  email: "",
  firstName: "",
  lastName: "",
};

interface IAuthContextType {
  auth: IContextType;
  setAuth: (value: IContextType) => void;
}

export const AuthContext = createContext<IAuthContextType | null>(null);
