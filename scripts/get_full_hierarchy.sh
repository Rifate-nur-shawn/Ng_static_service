#!/bin/bash

# Script to fetch complete Bangladesh location hierarchy
# Usage: ./scripts/get_full_hierarchy.sh

API_BASE="http://localhost:8080/api/v1/locations"
OUTPUT_FILE="location_hierarchy.json"

echo "🌍 Fetching complete Bangladesh location hierarchy..."
echo "This may take a moment..."
echo

# Start JSON structure
echo "{" > "$OUTPUT_FILE"
echo '  "divisions": [' >> "$OUTPUT_FILE"

# Get all divisions
divisions=$(curl -s "$API_BASE/divisions" | jq -r '.divisions[] | @json')

division_count=0
total_divisions=$(echo "$divisions" | wc -l)

echo "$divisions" | while IFS= read -r division; do
    division_count=$((division_count + 1))
    
    div_id=$(echo "$division" | jq -r '.id')
    div_name=$(echo "$division" | jq -r '.nameEn')
    
    echo "  📍 Processing: $div_name (ID: $div_id)"
    
    # Start division object
    echo "    {" >> "$OUTPUT_FILE"
    echo "      \"id\": \"$div_id\"," >> "$OUTPUT_FILE"
    echo "      \"nameEn\": \"$div_name\"," >> "$OUTPUT_FILE"
    echo "      \"districts\": [" >> "$OUTPUT_FILE"
    
    # Get districts for this division
    districts=$(curl -s "$API_BASE/divisions/$div_id/districts" | jq -r '.districts[] | @json')
    
    district_count=0
    total_districts=$(echo "$districts" | wc -l)
    
    echo "$districts" | while IFS= read -r district; do
        district_count=$((district_count + 1))
        
        dist_id=$(echo "$district" | jq -r '.id')
        dist_name=$(echo "$district" | jq -r '.nameEn')
        
        echo "     ├─ $dist_name (ID: $dist_id)"
        
        # Start district object
        echo "        {" >> "$OUTPUT_FILE"
        echo "          \"id\": \"$dist_id\"," >> "$OUTPUT_FILE"
        echo "          \"nameEn\": \"$dist_name\"," >> "$OUTPUT_FILE"
        echo "          \"upazilas\": [" >> "$OUTPUT_FILE"
        
        # Get upazilas for this district
        upazilas=$(curl -s "$API_BASE/districts/$dist_id/upazilas" | jq -r '.upazilas[] | @json')
        
        upazila_count=0
        total_upazilas=$(echo "$upazilas" | wc -l)
        
        echo "$upazilas" | while IFS= read -r upazila; do
            upazila_count=$((upazila_count + 1))
            
            up_id=$(echo "$upazila" | jq -r '.id')
            up_name=$(echo "$upazila" | jq -r '.nameEn')
            
            # Write upazila
            echo "            {" >> "$OUTPUT_FILE"
            echo "              \"id\": \"$up_id\"," >> "$OUTPUT_FILE"
            echo "              \"nameEn\": \"$up_name\"" >> "$OUTPUT_FILE"
            
            if [ $upazila_count -eq $total_upazilas ]; then
                echo "            }" >> "$OUTPUT_FILE"
            else
                echo "            }," >> "$OUTPUT_FILE"
            fi
        done
        
        # Close upazilas array and district object
        echo "          ]" >> "$OUTPUT_FILE"
        
        if [ $district_count -eq $total_districts ]; then
            echo "        }" >> "$OUTPUT_FILE"
        else
            echo "        }," >> "$OUTPUT_FILE"
        fi
    done
    
    # Close districts array and division object
    echo "      ]" >> "$OUTPUT_FILE"
    
    if [ $division_count -eq $total_divisions ]; then
        echo "    }" >> "$OUTPUT_FILE"
    else
        echo "    }," >> "$OUTPUT_FILE"
    fi
done

# Close JSON structure
echo "  ]" >> "$OUTPUT_FILE"
echo "}" >> "$OUTPUT_FILE"

echo
echo "✅ Complete hierarchy saved to: $OUTPUT_FILE"
echo
echo "📊 Summary:"
curl -s "$API_BASE/divisions" | jq -r '.divisions | length' | xargs echo "   Divisions:"
docker exec matrimonial_postgres psql -U postgres -d matrimonial_db -t -c "SELECT count(*) FROM districts;" | xargs echo "   Districts:"
docker exec matrimonial_postgres psql -U postgres -d matrimonial_db -t -c "SELECT count(*) FROM upazilas;" | xargs echo "   Upazilas:"
