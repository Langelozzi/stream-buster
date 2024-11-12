import React, { useCallback, useEffect, useState } from "react";
import { User } from "../../models/user";
import { CollapsibleListItem } from "../collapsible-list-item/CollapsibleListItem";
import { UsageStats } from "./UsageStats";
import { UserEndpointUsage } from "../../models/user_endpoint_usage";
import { getUserUsage } from "../../api/services/user.service";
import { sum } from "lodash";
import { TotalUsageProgress } from "./TotalUsageProgress";

interface UserUsageInfoProps {
    user: User;
    isAdmin: boolean;
}

export const UserUsageInfo: React.FC<UserUsageInfoProps> = ({ user, isAdmin }) => {
    const maxRequests: number = user.UserRoles[0].Role.MaxRequestCount;

    const [open, setOpen] = useState<boolean>(false);
    const [usage, setUsage] = useState<UserEndpointUsage[] | undefined>();
    const [requestCount, setRequestCount] = useState<number>(0);

    const handleToggle = () => {
        setOpen(prevOpen => !prevOpen);
    }

    // We need to refetch the usage for the user everytime this component reloads
    const fetchUsage = async () => {
        try {
            const fetchedUsage = await getUserUsage(user.ID);
            console.log(fetchedUsage);
            setUsage(fetchedUsage);
        } catch (e) {
            setUsage(undefined);
        }
    };

    const fetchAllEndpoints = async () => {
        // get all the endpoints so we can show all but set count to 0 for the endpoint
    }

    const calculateRequestCount = useCallback(() => {
        if (!usage) return;

        const count = sum(usage.map(endpointUsage => endpointUsage.RequestCount));
        setRequestCount(count);
    }, [usage]);

    useEffect(() => {
        fetchUsage();
    }, [user.ID]);

    useEffect(() => {
        calculateRequestCount();
    }, [calculateRequestCount]);

    return (
        <CollapsibleListItem
            label={`${user.FirstName} ${user.LastName}`}
            open={open}
            onToggle={handleToggle}
        >
            {requestCount && <TotalUsageProgress maxRequests={maxRequests} requestCount={requestCount} isAdmin={isAdmin} />}
            {usage && <UsageStats usage={usage} />}
        </CollapsibleListItem>
    )
}