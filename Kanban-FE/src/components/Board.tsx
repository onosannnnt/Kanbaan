import React, { useState, useEffect } from "react";
import { axiosInstance } from "../utils/axios";
import { useParams } from "react-router";

// const BOARD_ID = "1b01c7d7-3a41-4fa5-8c4a-962f9a666545"; // Example board ID

const Column = ({
  column,
  tasks,
  onAddTask,
  onAddAssign,
  onRenameTask,
  onDropTask,
  onRenameColumn,
  onDeleteColumn,
  onDeleteTask,
}) => {
  const [newTask, setNewTask] = useState("");
  const [newMembers, setNewMembers] = useState<Record<string, string>>({});
  const [isRenaming, setIsRenaming] = useState(false);
  const [newTitle, setNewTitle] = useState(column.Name);
  const [newTaskName, setNewTaskName] = useState<Record<string, string>>({});

  const handleDragOver = (e) => e.preventDefault();

  const handleDrop = (e) => {
    e.preventDefault();
    const taskData = JSON.parse(e.dataTransfer.getData("application/json"));
    onDropTask(taskData, column.ID);
  };

  const handleRename = () => {
    onRenameColumn(column.ID, newTitle);
    setIsRenaming(false);
  };

  const handleAddTask = () => {
    onAddTask(column.ID, newTask);
    setNewTask("");
  };
  const handleAddAssign = (taskID) => {
    onAddAssign(taskID, newMembers[taskID]);
    setNewMembers({});
  };
  const handleRenameTask = (taskID, newName) => {
    onRenameTask(taskID, newName);
    setNewTaskName({});
  };
  const handleDeleteTask = (taskID) => {
    onDeleteTask(taskID);
  };

  return (
    <div className="flex-1 p-4" onDragOver={handleDragOver} onDrop={handleDrop}>
      <div className="bg-gray-100 rounded-lg p-4 min-h-[200px]">
        <div className="flex justify-between items-center">
          {isRenaming ? (
            <input
              value={newTitle}
              onChange={(e) => setNewTitle(e.target.value)}
              onBlur={handleRename}
              onKeyDown={(e) => e.key === "Enter" && handleRename()}
              autoFocus
              className="border rounded px-2 py-1 w-full"
            />
          ) : (
            <h2
              className="text-xl font-semibold mb-2 cursor-pointer"
              onClick={() => setIsRenaming(true)}
            >
              {column.Name}
            </h2>
          )}
          <button
            onClick={() => onDeleteColumn(column.ID)}
            className="ml-2 text-red-500 hover:text-red-700"
          >
            &times;
          </button>
        </div>
        {tasks.map((task) => (
          <div className="mb-4 p-2 rounded bg-white" key={task.ID}>
            <div
              key={task.ID}
              draggable
              onDragStart={(e) =>
                e.dataTransfer.setData(
                  "application/json",
                  JSON.stringify({ ...task, fromColumn: column.ID })
                )
              }
              className="bg-white rounded p-2 shadow mb-2 cursor-grab"
            >
              {task.Name}
            </div>
            <div className="flex gap-2 mt-4">
              <input
                type="text"
                placeholder="New task"
                value={newMembers[task.ID]}
                onChange={(e) =>
                  setNewMembers({ ...newMembers, [task.ID]: e.target.value })
                }
                className="flex-1 px-2 py-1 border rounded"
              />
              <button
                onClick={() => handleAddAssign(task.ID)}
                disabled={!newMembers[task.ID]}
                className="px-4 py-1 bg-blue-600 text-white rounded disabled:opacity-50"
              >
                Assign
              </button>
            </div>
            <div className="text-sm text-gray-500 mt-2">
              <div>Assign To</div>
              <div className="flex flex-col gap-2">
                {task.Assignee?.map((member) => (
                  <div key={member.ID} className="text-sm">
                    {member.FirstName} {member.LastName}
                  </div>
                ))}
              </div>
            </div>
            <div className="flex gap-2 mt-4">
              <input
                type="text"
                placeholder="New task"
                value={newTaskName[task.ID] || ""}
                onChange={(e) =>
                  setNewTaskName({ ...newTaskName, [task.ID]: e.target.value })
                }
                className="flex-1 px-2 py-1 border rounded"
              />
              <button
                onClick={() => handleRenameTask(task.ID, newTaskName[task.ID])}
                disabled={!newTaskName[task.ID]}
                className="px-4 py-1 bg-blue-600 text-white rounded disabled:opacity-50"
              >
                Rename
              </button>
              <button
                onClick={() => handleDeleteTask(task.ID)}
                className="px-4 py-1 bg-red-600 text-white rounded disabled:opacity-50"
              >
                Delete
              </button>
            </div>
          </div>
        ))}
        <div className="flex gap-2 mt-4">
          <input
            type="text"
            placeholder="New task"
            value={newTask}
            onChange={(e) => setNewTask(e.target.value)}
            className="flex-1 px-2 py-1 border rounded"
          />
          <button
            onClick={handleAddTask}
            disabled={!newTask.trim()}
            className="px-4 py-1 bg-blue-600 text-white rounded disabled:opacity-50"
          >
            Add Task
          </button>
        </div>
      </div>
    </div>
  );
};

export default function KanbanBoard() {
  const BOARD_ID = useParams().id;
  const [columns, setColumns] = useState([]);
  const [tasks, setTasks] = useState([]);
  const [newColumnName, setNewColumnName] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [colRes, taskRes] = await Promise.all([
          axiosInstance.get(`/columns/board/${BOARD_ID}`),
          axiosInstance.get(`/tasks/`),
        ]);
        setColumns(colRes.data.data);
        setTasks(taskRes.data.data);
      } catch (error) {
        console.error("Failed to fetch data:", error);
      }
    };
    fetchData();
  }, []);

  const reloadTasks = async () => {
    const res = await axiosInstance.get(`/tasks/`);
    setTasks(res.data.data);
  };

  const onAddTask = async (columnId, name) => {
    try {
      await axiosInstance.post(`/tasks/`, {
        name,
        description: "",
        column_id: columnId,
      });
      reloadTasks();
    } catch (err) {
      console.error(err);
    }
  };

  const onAddAssign = async (taskID: string, email: string) => {
    try {
      const response = await axiosInstance.post(`/tasks/${taskID}/assigns/`, {
        assignee_id: [email],
      });
      console.log("response", response);
      if (response.status === 200) {
        window.location.reload();
      }
    } catch (err) {
      console.error(err);
    }
  };

  const onRenameTask = async (taskID, newName) => {
    try {
      await axiosInstance.put(`/tasks/${taskID}`, {
        name: newName,
      });
      reloadTasks();
    } catch (err) {
      console.error(err);
    }
  };

  const onDeleteTask = async (taskID) => {
    try {
      await axiosInstance.delete(`/tasks/${taskID}`);
      reloadTasks();
    } catch (err) {
      console.error(err);
    }
  };

  const onDropTask = async (task, targetColumnId) => {
    try {
      await axiosInstance.put(`/tasks/${task.ID}`, {
        column_id: targetColumnId,
      });
      reloadTasks();
    } catch (err) {
      console.error(err);
    }
  };

  const addColumn = async () => {
    if (!newColumnName.trim()) return;
    try {
      await axiosInstance.post(`/columns`, {
        name: newColumnName,
        board_id: BOARD_ID,
      });
      const res = await axiosInstance.get(`/columns/board/${BOARD_ID}`);
      setColumns(res.data.data);
      setNewColumnName("");
    } catch (err) {
      console.error(err);
    }
  };

  const onRenameColumn = async (columnId, newName) => {
    try {
      await axiosInstance.put(`/columns/${columnId}`, {
        name: newName,
      });
      const res = await axiosInstance.get(`/columns/board/${BOARD_ID}`);
      setColumns(res.data.data);
    } catch (err) {
      console.error(err);
    }
  };

  const onDeleteColumn = async (columnId) => {
    try {
      await axiosInstance.delete(`/columns/${columnId}`);
      setColumns((prev) => prev.filter((col) => col.ID !== columnId));
      setTasks((prev) => prev.filter((task) => task.column_id !== columnId));
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold text-center mb-6">Kanban Board</h1>
      <div className="flex gap-2 mb-4">
        <input
          value={newColumnName}
          onChange={(e) => setNewColumnName(e.target.value)}
          placeholder="New column name"
          className="flex-1 px-2 py-1 border rounded"
        />
        <button
          onClick={addColumn}
          className="px-4 py-1 bg-green-600 text-white rounded"
        >
          Add Column
        </button>
      </div>
      <div className="flex flex-row gap-4 flex-wrap">
        {columns.map((col) => (
          <Column
            key={col.ID}
            column={col}
            tasks={tasks.filter((t) => t.column_id === col.ID)}
            onAddTask={onAddTask}
            onAddAssign={onAddAssign}
            onRenameTask={onRenameTask}
            onDropTask={onDropTask}
            onRenameColumn={onRenameColumn}
            onDeleteColumn={onDeleteColumn}
            onDeleteTask={onDeleteTask}
          />
        ))}
      </div>
    </div>
  );
}
