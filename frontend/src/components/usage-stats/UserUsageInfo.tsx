import React, { useCallback, useEffect, useState } from "react";
import { User } from "../../models/user";
import { ExpandableCard } from "../explandable-card/ExpandableCard";
import { UsageStats } from "./UsageStats";
import { UserEndpointUsage } from "../../models/user_endpoint_usage";
import { getUserUsage } from "../../api/services/user.service";
import { sum } from "lodash";
import { TotalUsageProgress } from "./TotalUsageProgress";
import { getAllEndpoints } from "../../api/services/endpoint.service";
import { Endpoint } from "../../models/endpoint";

interface UserUsageInfoProps {
    user: User;
    isAdmin: boolean;
}

export const UserUsageInfo: React.FC<UserUsageInfoProps> = ({ user, isAdmin }) => {
    const maxRequests: number = user.UserRoles[0].Role.MaxRequestCount;
    const trackedEndpoints = [
        '/api/v1/search/multi',
        '/api/v1/cdn/movie/:tmdbId',
        '/api/v1/cdn/tv/:tmdbId/:seasonNum/:episodeNum',
        '/api/v1/tv/:id',
        '/api/v1/tv/:id/season/:seasonNum/episodes',
        '/api/v1/movie/:id',
        '/api/v1/llm/ask-query'
    ]

    const [usage, setUsage] = useState<UserEndpointUsage[] | undefined>();
    const [endpoints, setEndpoints] = useState<Endpoint[]>();
    const [filteredEndpoints, setFilteredEndpoints] = useState<Endpoint[]>();
    const [requestCount, setRequestCount] = useState<number>(0);

    // We need to refetch the usage for the user everytime this component reloads
    const fetchUsage = async () => {
        try {
            const fetchedUsage = await getUserUsage(user.ID);
            setUsage(fetchedUsage);
        } catch (e) {
            setUsage(undefined);
        }
    };

    const fetchAllEndpoints = async () => {
        // get all the endpoints so we can show all but set count to 0 for the endpoint
        try {
            const fetchedEndpoints = await getAllEndpoints();
            setEndpoints(fetchedEndpoints);
        } catch (e) {
            console.error(e);
        }
    }

    const calculateRequestCount = useCallback(() => {
        if (!usage) return;

        const count = sum(usage.map(endpointUsage => endpointUsage.RequestCount));
        setRequestCount(count);
    }, [usage]);

    const filterEndpoints = useCallback(() => {
        const filtered = endpoints?.filter(endpoint => trackedEndpoints.includes(endpoint.Path));
        setFilteredEndpoints(filtered);
    }, [endpoints])

    useEffect(() => {
        fetchAllEndpoints();
    }, [])

    useEffect(() => {
        fetchUsage();
    }, [user.ID]);

    useEffect(() => {
        calculateRequestCount();
    }, [calculateRequestCount]);

    useEffect(() => {
        filterEndpoints();
    }, [filterEndpoints])

    return (
        <ExpandableCard
            headerContent={
                <TotalUsageProgress maxRequests={maxRequests} requestCount={requestCount} isAdmin={isAdmin} />
            }
            expandedContent={
                usage && filteredEndpoints && <UsageStats usage={usage} endpoints={filteredEndpoints} />
            }
        />
    )
}