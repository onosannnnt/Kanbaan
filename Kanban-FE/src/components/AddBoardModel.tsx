import React, { useState } from "react";
import { axiosInstance } from "../utils/axios";

interface AddBoardModelProps {
  setAddModel: (model: boolean) => void;
}

const AddBoardModel: React.FC<AddBoardModelProps> = ({ setAddModel }) => {
  const [boardName, setBoardName] = useState<string>("");
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setBoardName(e.target.value);
  };
  const handleAddBoard = async () => {
    if (!boardName) {
      return;
    }
    try {
      const response = await axiosInstance.post("/boards", {
        name: boardName,
      });
      if (response.status === 200) {
        setAddModel(false);
      }
    } catch (error) {
      console.error("Error adding board:", error);
    }
  };
  return (
    <>
      <form className="fixed inset-0 flex items-center justify-center">
        <div className="bg-white p-4 rounded shadow-md">
          <h2 className="text-lg font-bold mb-4">Add Board</h2>
          <input
            type="text"
            placeholder="Board Name"
            onChange={handleInputChange}
            className="border border-gray-300 p-2 rounded mb-4 w-full"
          />
          <div className="grid grid-cols-2 items-center gap-2">
            <button
              onClick={() => {
                setAddModel(false);
              }}
              className="bg-red-500 text-white p-2 rounded hover:bg-red-600 transition duration-200"
            >
              Cancel
            </button>
            <button
              onClick={handleAddBoard}
              className="bg-green-500 text-white p-2 rounded hover:bg-green-600 transition duration-200"
            >
              Save
            </button>
          </div>
        </div>
      </form>
    </>
  );
};

export default AddBoardModel;
