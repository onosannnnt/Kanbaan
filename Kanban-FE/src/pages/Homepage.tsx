import ColabBoard from "../components/ColabBoard";
import MyBoard from "../components/MyBoard";

const Homepage = () => {
  return (
    <main className="w-full h-fit bg-slate-200">
      <MyBoard></MyBoard>
      <ColabBoard></ColabBoard>
    </main>
  );
};

export default Homepage;
