# MME

## The Media Metadata Extractor (MME Service)
The Concept: A service that accepts file upload payloads (or links to raw media) and extracts useful metadata from them.
Input (Upstream): A JSON payload containing a link to an image file or raw image bytes.
Transformation (mme service): It parses the image, detects the file format, dimensions (width/height), file size, and extracts any EXIF data (like GPS coordinates or camera model).
Output (Downstream): A clean, structured JSON object sent to an image-optimization or storage service.
Example Output: {"format": "png", "width": 1920, "height": 1080, "size_bytes": 204800}


# 1. The Data Flow Architecture
To visualize how data flows through this service to the downstream service (like a Database or an Image Optimization service).
2. Mapping the Logic to Your Folder Structure
Here is how we will distribute the code for the Media Metadata Extractor across our new directories:

api/protobuf/service.proto
Define the contract for our service. Even if we start with standard HTTP/JSON, defining the interface here keeps our design clean.
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

# Architecture Model

To model this, we use elements from the Application Layer alongside a couple of external Business Layer actors representing the boundary triggers.
## Active Structure Elements
Actor: Upstream Client / API User – The user or automated script kicking off the process.
Application Interface: HTTP REST API (:8080/extract) – The exposed boundary interface of our application.
Application Component: Media Metadata Extractor Service – The primary software component container (our microservice).
Application Interface: Downstream Client HTTP Interface – The outgoing connection client.
Application Component: Downstream Mock Service (:8081/receiver) – The external consuming application.
## Behavior Elements
Application Function: Image Fetching & Extraction – The core internal logic handled by our ExtractorService.
Application Function: Metadata Forwarding – The dispatch action handled by our DownstreamClient repository.
## Passive Structure Elements (Data)
Application Event/Data Object: ExtractRequest JSON – The inbound data payload container (image_url).
Application Data Object: ImageMetadata JSON – The processed core domain payload (format, width, height, size_bytes).



