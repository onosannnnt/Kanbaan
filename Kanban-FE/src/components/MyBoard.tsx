import { CiCirclePlus } from "react-icons/ci";
import BoardCard from "./BoardCard";
import { axiosInstance } from "../utils/axios";
import { useEffect, useState } from "react";
import Loading from "./Loading";
import AddBoardModel from "./AddBoardModel";

interface IBoard {
  ID: string;
  Name: string;
  CreatedAt: string;
}

const ColabBoard = () => {
  const [board, setBoard] = useState<IBoard[]>([]);
  const [addModel, setAddModel] = useState(false);

  const [loading, setLoading] = useState(true);

  const fetchBoard = async () => {
    try {
      const response = await axiosInstance.get("/users/me/boards/");
      if (response.status === 200) {
        setBoard(response.data.data);
      }
      setLoading(false);
    } catch (error) {
      setLoading(false);
      console.error("Error fetching board data:", error);
    }
  };
  useEffect(() => {
    fetchBoard();
  }, []);

  if (loading) {
    return (
      <div className="flex items-center justify-center w-full h-screen">
        <Loading />
      </div>
    );
  }

  return (
    <main className="w-full h-fit flex flex-col p-16">
      <div className="flex items-center gap-4">
        <h1 className="font-bold text-4xl">My board</h1>
        <div
          onClick={
            addModel ? () => setAddModel(false) : () => setAddModel(true)
          }
        >
          <CiCirclePlus className="text-4xl hover:text-green-700 transition duration-200" />
        </div>
      </div>

      <div className="flex flex-col gap-4 mt-8">
        <div className="w-full h-fit grid grid-cols-3 px-16 py-4 gap-8 bg-slate-300 rounded-lg">
          {board.length === 0 ? (
            <div className="flex items-center justify-center col-span-3 text-4xl">
              No board
            </div>
          ) : (
            board.map((item: IBoard) => (
              <BoardCard
                key={item.ID}
                ID={item.ID}
                title={item.Name}
                createdAt={item.CreatedAt}
                type={"myboard"}
              />
            ))
          )}
        </div>
        {addModel && <AddBoardModel setAddModel={setAddModel}></AddBoardModel>}
      </div>
    </main>
  );
};

export default ColabBoard;
