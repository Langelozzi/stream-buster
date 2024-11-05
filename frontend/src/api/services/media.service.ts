import instance from "../axios";

export const getMedia = (id: number) => {
    try {
        const res = instance.get("/media/", { params: { id: id } })
        console.log(res)
        return res
    } catch (error) {
        return error
    }
}
