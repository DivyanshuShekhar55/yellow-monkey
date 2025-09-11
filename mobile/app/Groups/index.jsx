import React from 'react';
import {
    View,
    Text,
    FlatList,
    TouchableOpacity,
    Image,
    StyleSheet,
    SafeAreaView,
    StatusBar,
} from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { useRouter } from "expo-router"

const SAMPLE_GROUPS = [
    {
        id: '1',
        name: 'Fashion Club',
        image: 'https://images.unsplash.com/photo-1441986300917-64674bd600d8?w=400',
        distance: '7 km',
        mutualFriends: [
            'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=50',
            'https://images.unsplash.com/photo-1494790108755-2616b612b29c?w=50',
            'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=50',
        ],
        tags: ['Fashion', 'Style', 'Trends'],
        description: 'A community for fashion enthusiasts to share styles and trends.',
        status: 'join'
    },
    {
        id: '2',
        name: 'Skateboarding in City Park',
        image: 'https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400',
        distance: '2 km',
        mutualFriends: [
            'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=50',
            'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=50',
        ],
        tags: ['Skating', 'Sports', 'Park'],
        description: 'Join us for skateboarding sessions at the city park every weekend.',
        status: 'requested'
    },
    {
        id: '3',
        name: 'Tech Innovators',
        image: 'https://images.unsplash.com/photo-1519389950473-47ba0277781c?w=400',
        distance: '12 km',
        mutualFriends: [
            'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=50',
            'https://images.unsplash.com/photo-1494790108755-2616b612b29c?w=50',
            'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=50',
            'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=50',
        ],
        tags: ['Tech', 'Innovation', 'Startup'],
        description: 'Connect with tech enthusiasts and innovators in your area.',
        status: 'open'
    },
];

export default function GroupListScreen() {
    const router = useRouter();
    const renderGroupItem = ({ item }) => (
        <TouchableOpacity
            style={styles.groupCard}
            onPress={() => {router.push({
                pathname: "/Groups/GroupDetails",
                params: { group_str: JSON.stringify(item) }
            })
            console.log(JSON.stringify(item))
        }
        }
        >
            <View style={styles.groupHeader}>
                <Image source={{ uri: item.image }} style={styles.groupImage} />
                <View style={styles.groupInfo}>
                    <Text style={styles.groupName}>{item.name}</Text>
                    <View style={styles.locationRow}>
                        <Ionicons name="location-outline" size={16} color="#888" />
                        <Text style={styles.distance}>{item.distance}</Text>
                    </View>
                </View>
            </View>

            <View style={styles.mutualsRow}>
                <View style={styles.mutualAvatars}>
                    {item.mutualFriends.slice(0, 4).map((avatar, index) => (
                        <Image
                            key={index}
                            source={{ uri: avatar }}
                            style={[styles.mutualAvatar, { marginLeft: index > 0 ? -8 : 0 }]}
                        />
                    ))}
                </View>
            </View>

            <View style={styles.tagsRow}>
                {item.tags.map((tag, index) => (
                    <View key={index} style={styles.tag}>
                        <Text style={styles.tagText}>{tag}</Text>
                    </View>
                ))}
            </View>
        </TouchableOpacity>
    );

    return (
        <SafeAreaView style={styles.container}>
            <StatusBar barStyle="light-content" backgroundColor="#1a1a1a" />
            <FlatList
                data={SAMPLE_GROUPS}
                renderItem={renderGroupItem}
                keyExtractor={(item) => item.id}
                contentContainerStyle={styles.listContainer}
                showsVerticalScrollIndicator={false}
            />
        </SafeAreaView>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#1a1a1a',
    },
    listContainer: {
        padding: 16,
    },
    groupCard: {
        backgroundColor: '#2a2a2a',
        borderRadius: 16,
        padding: 16,
        marginBottom: 16,
        borderWidth: 1,
        borderColor: '#333',
    },
    groupHeader: {
        flexDirection: 'row',
        alignItems: 'center',
        marginBottom: 12,
    },
    groupImage: {
        width: 60,
        height: 60,
        borderRadius: 30,
        marginRight: 16,
    },
    groupInfo: {
        flex: 1,
    },
    groupName: {
        fontSize: 18,
        fontWeight: '600',
        color: '#fff',
        marginBottom: 4,
    },
    locationRow: {
        flexDirection: 'row',
        alignItems: 'center',
    },
    distance: {
        fontSize: 14,
        color: '#888',
        marginLeft: 4,
    },
    mutualsRow: {
        flexDirection: 'row',
        alignItems: 'center',
        marginBottom: 12,
    },
    mutualAvatars: {
        flexDirection: 'row',
        alignItems: 'center',
    },
    mutualAvatar: {
        width: 32,
        height: 32,
        borderRadius: 16,
        borderWidth: 2,
        borderColor: '#2a2a2a',
    },
    tagsRow: {
        flexDirection: 'row',
        flexWrap: 'wrap',
    },
    tag: {
        backgroundColor: '#333',
        borderRadius: 20,
        paddingHorizontal: 12,
        paddingVertical: 6,
        marginRight: 8,
        marginBottom: 4,
    },
    tagText: {
        fontSize: 12,
        color: '#fff',
        fontWeight: '500',
    },
});