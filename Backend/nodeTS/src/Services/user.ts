import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

export default new class UserServices {
    async GetUsers() {
        try {
            const user = await prisma.users.findMany({
                select: {
                    id: true,
                    name: true,
                    email: true,
                    phoneNumber: true,
                    role: true,
                    createdAt: true,
                    updatedAt: true
                }
            })
            return {
                status: "Success",
                messages: "Success Get All User",
                data: user
            }
        } catch (error) {
            throw error
        }
    }

    async GetUserById(id: number) {
        try {
            const user = await prisma.users.findFirst({
                where: {
                    id: id
                },
                select: {
                    id: true,
                    name: true,
                    email: true,
                    phoneNumber: true,
                    role: true,
                    createdAt: true,
                    updatedAt: true
                }
            })
            
            if(!user) {
                return {
                    status: "Failed",
                    messages: `id = ${id} not found !`,
                }
            } else {
                return {
                    status: "Success",
                    messages: "Success Get User With Id = " + id,
                    data: user
                }
            }
        } catch (error) {
            throw error
        }
    }

    async UpdateUser(id: number, data: object) {
        try {

            const user = await prisma.users.update({
                where: {
                    id: id
                },
                data,
            })

            if(!user) {
                return {
                    status: "Failed",
                    messages: `User with id ${id}, Not Found`,
                }
            } else {
                return {
                    status: "Success",
                    messages: "Success Update User",
                    data: user
                }
            }

        } catch (error) {
            throw error
        }
    }

    async DeleteUsers(id: number) {
        try {
            const user = await prisma.users.delete({
                where: {
                    id: id
                }
            })

            if(!user) {
                return {
                    status: 'Failed',
                    messages: `Failed delete user id = ${id}`
                }
            } else {
                return {
                    status: 'Success',
                    messages: `Success delete user ${user.email}`
                }
            }
            
        } catch (error) {
            throw error
        }
    }
}