import instance from "../axios";

export const getContentExists = async (tmdbId: number, isTV: boolean) => {
    try {
        const res = await instance.get(
            `/cdn/${tmdbId}/exists`,
            {
                params: {
                    isTV: isTV
                }
            }
        );

        return res.data.exists;
    } catch (error) {
        return error
    }
}