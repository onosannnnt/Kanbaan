import type React from "react";
import { useState } from "react";
import { axiosInstance } from "../utils/axios";

export interface IResponse {
  data: IData;
  message: string;
  status: number;
}

export interface IData {
  CreatedAt: string;
  DeletedAt: null;
  ID: string;
  Members: IMember[] | null;
  Name: string;
  Owner: IOwner;
  OwnerID: string;
  UpdatedAt: string;
}

export interface IMember {
  CreatedAt: string;
  DeletedAt: null;
  Email: string;
  FirstName: string;
  ID: string;
  LastName: string;
  Password: string;
  UpdatedAt: string;
}

export interface IOwner {
  CreatedAt: string;
  DeletedAt: null;
  Email: string;
  FirstName: string;
  ID: string;
  LastName: string;
  Password: string;
  UpdatedAt: string;
}

interface ShareModelProps {
  setModel: (model: boolean) => void;
  board: IData;
}

const ShareModel: React.FC<ShareModelProps> = ({ setModel, board }) => {
  const [email, setEmail] = useState<string>();
  const handleInvite = async (e: React.FormEvent<HTMLFormElement>) => {
    e.stopPropagation();
    try {
      e.preventDefault();
      const response = await axiosInstance.post(`/boards/${board.ID}/invites`, {
        members: [email],
      });
      if (response.status === 200) {
        setModel(false);
        window.location.reload();
      }
    } catch (error) {
      console.error("Error inviting user:", error);
    }
  };
  return (
    <form
      onSubmit={handleInvite}
      className="fixed inset-0 flex items-center justify-center"
    >
      <div className="w-1/3 h-1/3 bg-white rounded-lg shadow-lg flex flex-col items-center justify-center">
        <h1 className="text-2xl font-bold mb-4">Share Board</h1>
        <input
          type="email"
          onChange={(e) => setEmail(e.target.value)}
          onClick={(e) => e.stopPropagation()}
          placeholder="Enter email to share"
          className="p-2 w-3/4 border-b-2 border-gray-300 rounded mb-4"
        />
        <div className="flex items-center gap-2">
          <button
            onClick={(e: React.MouseEvent<HTMLButtonElement>) => {
              e.stopPropagation();
              setModel(false);
            }}
            className="bg-red-500 text-white p-2 rounded hover:bg-red-600 transition duration-200"
          >
            Cancel
          </button>
          <button className="bg-green-500 text-white p-2 rounded hover:bg-green-600 transition duration-200">
            Share
          </button>
        </div>

        <div className="flex flex-col items-center justify-center mt-4 w-full px-4">
          <h1 className="text-2xl font-bold">Members</h1>
          <div className="flex flex-col items-center justify-center mt-4 w-full px-4 gap-2">
            {board?.Members?.map((item: IMember) => (
              <div
                key={item.ID}
                className="grid grid-cols-2 items-center gap-2 w-full"
              >
                <h1>
                  {item.FirstName} {item.LastName}
                </h1>
              </div>
            ))}
          </div>
        </div>
      </div>
    </form>
  );
};

export default ShareModel;
