import { Media } from "../../models/media";
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

export const createMedia = async (media: Media) => {
    try {
        const res = instance.post("/media/create", media);
        console.log('res', res);
        return res
    } catch (error) {
        throw error
    }
}
