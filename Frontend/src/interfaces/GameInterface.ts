export interface GameState {
  id: number | null;
  name: string | null;
  description: string | null;
  icon: string | null;
  types: string[] | null;
}

export interface GameData {
  data: GameState[] | null;
  status: "idle" | "loading" | "success" | "error";
}
