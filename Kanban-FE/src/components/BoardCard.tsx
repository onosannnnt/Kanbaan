import { useEffect, useState } from "react";
import ShareModel, { type IData } from "./ShareModel";
import { axiosInstance } from "../utils/axios";

interface BoardCardProps {
  ID: string;
  title: string;
  createdAt: string;
  type?: string;
}

const BoardCard = (props: BoardCardProps) => {
  const { ID, createdAt } = props;
  const [model, setModel] = useState(false);
  const [board, setBoard] = useState<IData>();

  const onOpenModel = (e: React.MouseEvent) => {
    e.stopPropagation();
    setModel(true);
  };
  const fetchBoard = async () => {
    try {
      const response = await axiosInstance.get(`/boards/${ID}`);
      if (response.status === 200) {
        setBoard(response.data.data);
      }
    } catch (error) {
      console.error("Error fetching board data:", error);
    }
  };
  const onChangeTitle = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newTitle = e.target.value;
    setBoard((prev) => {
      if (prev) {
        return { ...prev, Name: newTitle };
      }
      return prev;
    });
  };

  const onUpdateBoard = async (e: React.MouseEvent<HTMLElement>) => {
    e.stopPropagation();
    try {
      const response = await axiosInstance.put(`/boards/${ID}`, {
        Name: board?.Name,
      });
      if (response.status === 200) {
        console.log("Board updated successfully");
      }
    } catch (error) {
      console.error("Error updating board:", error);
    }
  };
  const handleBoardClick = () => {
    window.location.href = `/board/${ID}`;
  };
  const onDeleteBoard = async () => {
    try {
      const response = await axiosInstance.delete(`/boards/${ID}`);
      if (response.status === 204) {
        window.location.reload();
      }
    } catch (error) {
      console.error("Error deleting board:", error);
    }
  };
  const date = new Date(createdAt);

  useEffect(() => {
    fetchBoard();
  }, []);
  return (
    <main
      onClick={handleBoardClick}
      className={`flex flex-col gap-2 p-2 w-3/4 h-full cursor-pointer bg-slate-200 rounded-lg`}
    >
      {props.type === "myboard" ? (
        <input
          type="text"
          onChange={onChangeTitle}
          onClick={(e) => e.stopPropagation()}
          value={board?.Name}
          className="text-2xl font-bold static z-10"
        />
      ) : (
        <div className="text-2xl font-bold">{board?.Name}</div>
      )}
      <div className="flex items-center gap-2 text-sm text-gray-500">
        {date.getDate()}/{date.getMonth()}/{date.getFullYear()}
      </div>
      {props.type === "myboard" && (
        <div className="flex items-center gap-2 text-sm">
          <div
            onClick={(e) => {
              onUpdateBoard(e);
            }}
            className="text-blue-500 hover:text-blue-600 transition duration-200 cursor-pointer"
          >
            Save
          </div>
          <div
            onClick={onOpenModel}
            className="text-blue-500 hover:text-blue-600 transition duration-200 cursor-pointer"
          >
            Share
          </div>
          <div
            onClick={onDeleteBoard}
            className="text-red-500 hover:text-red-600 transition duration-200 cursor-pointer"
          >
            Delete
          </div>
        </div>
      )}

      {model && board && (
        <ShareModel setModel={setModel} board={board}></ShareModel>
      )}
    </main>
  );
};

export default BoardCard;
