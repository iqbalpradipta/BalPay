import { Prisma, PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

export default new (class GameServices {
  async CreateGame(data: Prisma.GamesCreateInput) {
    try {
      await prisma.games.create({
        data,
      });

      return {
        status: "Success",
        messages: "Success Create data",
        data,
      };
    } catch (error) {
      throw error;
    }
  }

  async GetGames() {
    try {
      const game = await prisma.games.findMany();

      return {
        status: "Success",
        messages: "Success Get data",
        data: game,
      };
    } catch (error) {
      throw error;
    }
  }

  async GetGameById(id: number) {
    try {
      const game = await prisma.games.findFirst({
        where: {
          id,
        },
      });

      if (!game) {
        return {
          status: "Failed",
          messages: "Failed get data by Id",
        };
      } else {
        return {
          status: "Success",
          messages: "Success get data by Id",
          data: game,
        };
      }
    } catch (error) {
      throw error;
    }
  }

  async UpdateGame(id: number, data: Prisma.GamesCreateInput) {
    try {
      const game = await prisma.games.update({
        where: {
          id: id,
        },
        data,
      });

      if (!game) {
        return {
          status: "Failed",
          messages: "Failed update data",
        };
      } else {
        return {
          status: "Success",
          messages: "Success update data",
          data: game,
        };
      }
    } catch (error) {
      throw error;
    }
  }

  async DeleteGame(id: number) {
    try {
      const game = await prisma.games.delete({
        where: {
          id: id,
        },
      });

      if (!game) {
        return {
          status: "Failed",
          messages: "Failed delete data",
        };
      } else {
        return {
          status: "Success",
          messages: `Success delete data ${game.name}`,
          data: game,
        };
      }
    } catch (error) {
      throw error;
    }
  }
})();
