import { Prisma, PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

export default new class GameServices {
    async CreateGame(data: Prisma.GamesCreateInput) {
        try {
            await prisma.games.create({
                data,
            })

            return {
                status: "Success",
                messages: "Success Create data",
                data,
            }
        } catch (error) {
            throw error
        }
    }

    async GetGames() {
        try {
            const game = await prisma.games.findMany()

            return {
                status: "Success",
                messages: "Success Get data",
                data: game,
            }
        } catch (error) {
            throw error
        }
    } 

    async GetGameById(id: number) {
        try {
            const game = await prisma.games.findFirst({
                where: {
                    id,
                }
            })

            return {
                status: "Success",
                messages: "Success get data by Id",
                data: game,
            }
        } catch (error) {
            return {
                status: "Failed",
                messages: new Error
            }
        }
    }

    async UpdateGame(id: number, data: Prisma.GamesCreateInput) {
        try {
            const game = await prisma.games.update({
                where: {
                    id: id,
                },
                data,
            })

            return {
                status: "Success",
                messages: "Success update data",
                data: game,
            }
        } catch (error) {
            return {
                status: "Failed",
                messages: new Error
            }
        }
    }

    async DeleteGame(id: number) {
        try {
            const game = await prisma.games.delete({
                where: {
                    id: id
                }
            })

            return {
                status: "Success",
                messages: "Success delete data",
            }
        } catch (error) {
            return {
                status: "Failed",
                messages: new Error
            }
        }
    }
}