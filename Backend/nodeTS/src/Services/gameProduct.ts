import { Prisma, PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

export default new class GameProductServices{
    async CreateGameProduct(data: Prisma.GameProductCreateInput) {
        try {
            const product = await prisma.gameProduct.create({data})

            return {
                status: 'Success',
                messages: 'Success Create Game Product',
                data,
            }
        } catch (error) {
            throw error
        }
    }

    async GetGameProduct() {
        try {
            const product = await prisma.gameProduct.findMany({
                select: {
                    name: true,
                    price: true,
                    image: true,
                    type: true,
                    games: true,
                    gameId: true,
                    createdAt: true,
                    updatedAt: true
                }
            })
            
            return {
                status: 'Success',
                messages: 'Success Get Game Product',
                data: product
            }
        } catch (error) {
            throw error
        }
    }

    async GetGameProductById(id: number) {
        try {
            const product = await prisma.gameProduct.findFirst({
                where: {
                    id: id
                },
                select: {
                    name: true,
                    price: true,
                    image: true,
                    type: true,
                    games: true,
                    gameId: true,
                    createdAt: true,
                    updatedAt: true
                }
            })
            
            if(!product) {
                return {
                    status: 'Failed',
                    messages: `id: ${id} not found !`,
                }
            } else {
                return {
                    status: 'Success',
                    messages: 'Success Get Game Product with id: '+ id,
                    data: product
                }
            }
        } catch (error) {
            throw error
        }
    }

    async UpdateGameProduct(id: number, data: Prisma.GameProductCreateInput) {
        try {
            const product = await prisma.gameProduct.update({
                where: {
                    id
                },
                data,
            })
            
            if(product.id !== id) {
                return {
                    status: 'Failed',
                    messages: `id: ${id} not found !`,
                }
            } else {
                return {
                    status: 'Success',
                    messages: 'Success Update Game Product',
                    data: product
                }
            }
        } catch (error) {
            throw error
        }
    }

    async DeleteGameProduct(id: number) {
        try {
            const product = await prisma.gameProduct.delete({where: {id}})
            
            if(!product) {
                return {
                    status: 'Failed',
                    messages: `id: ${id} not found !`,
                }
            } else {
                return {
                    status: 'Success',
                    messages: 'Success Delete Game Product',
                    data: product
                }
            }
        } catch (error) {
            throw error
        }
    }
}