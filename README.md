# MME

## The Media Metadata Extractor (MME Service)
The Concept: A service that accepts file upload payloads (or links to raw media) and extracts useful metadata from them.
Input (Upstream): A JSON payload containing a link to an image file or raw image bytes.
Transformation (Your Service): It parses the image, detects the file format, dimensions (width/height), file size, and extracts any EXIF data (like GPS coordinates or camera model).
Output (Downstream): A clean, structured JSON object sent to an image-optimization or storage service.
Example Output: {"format": "png", "width": 1920, "height": 1080, "size_bytes": 204800}


# 1. The Data Flow Architecture
Before writing code, it helps to visualize how data flows through this service to the downstream service (like a Database or an Image Optimization service).
2. Mapping the Logic to Your Folder Structure
Here is how you will distribute the code for the Media Metadata Extractor across your new directories:

api/protobuf/service.proto
Define the contract for your service. Even if you start with standard HTTP/JSON, defining the interface here keeps your design clean.
Protocol Buffers


syntax = "proto3";
package metadata;

service MetadataExtractor {
    rpc ExtractFromUrl (ExtractRequest) returns (ExtractResponse);
}

message ExtractRequest {
    string image_url = 1;
}

message ExtractResponse {
    string format = 1;
    int32 width = 2;
    int32 height = 3;
    int64 size_bytes = 4;
}




