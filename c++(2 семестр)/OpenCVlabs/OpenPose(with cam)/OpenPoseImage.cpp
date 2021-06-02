#include <opencv2/dnn.hpp>
#include <opencv2/imgproc.hpp>
#include <opencv2/highgui.hpp>
#include <iostream>

using namespace std;
using namespace cv;
using namespace cv::dnn;

#define MPI
// rapidjson/example/simpledom/simpledom.cpp`
#include "rapidjson/document.h"
#include "rapidjson/writer.h"
#include "rapidjson/stringbuffer.h"
#include <iostream>
using namespace rapidjson;
#ifdef MPI
const int POSE_PAIRS[14][2] = 
{   
    {0,1}, {1,2}, {2,3},
    {3,4}, {1,5}, {5,6},
    {6,7}, {1,14}, {14,8}, {8,9},
    {9,10}, {14,11}, {11,12}, {12,13}
};

string protoFile = "pose/mpi/pose_deploy_linevec_faster_4_stages.prototxt";
string weightsFile = "pose/mpi/pose_iter_160000.caffemodel";

int nPoints = 15;
#endif

#ifdef COCO
const int POSE_PAIRS[17][2] = 
{   
    {1,2}, {1,5}, {2,3},
    {3,4}, {5,6}, {6,7},
    {1,8}, {8,9}, {9,10},
    {1,11}, {11,12}, {12,13},
    {1,0}, {0,14},
    {14,16}, {0,15}, {15,17}
};

string protoFile = "pose/coco/pose_deploy_linevec.prototxt";
string weightsFile = "pose/coco/pose_iter_440000.caffemodel";

int nPoints = 18;
#endif
void serializeToJson( const std::vector<Point> & myObjs) {
    rapidjson::Document jsonDoc;
    jsonDoc.SetArray();
    rapidjson::Document::AllocatorType& allocator = jsonDoc.GetAllocator();
    rapidjson::Value koorsArray(rapidjson::kArrayType);
    for (auto & koord : myObjs) {
        rapidjson::Value objKoord;
        objKoord.SetObject();
        objKoord.AddMember("x", koord.x, allocator);
        objKoord.AddMember("y", koord.y, allocator);
        koorsArray.PushBack(objKoord, allocator);
    }
    rapidjson::Value objValue;
    objValue.SetObject();
    objValue.AddMember("Coordinat", koorsArray, allocator);
    jsonDoc.PushBack(objValue, allocator);

    rapidjson::StringBuffer strbuf;
    rapidjson::Writer<rapidjson::StringBuffer> writer(strbuf);
    jsonDoc.Accept(writer);
    cout << strbuf.GetString() << endl;
}
int main(int argc, char **argv)
{

    
    string device = "cpu";



    int inWidth = 368;
    int inHeight = 368;
    float thresh = 0.1;



    VideoCapture cam{0};

    if (cam.isOpened())
    {
        String window = "Mirrow";

        namedWindow(window);

        auto run = true;
        Mat img;
        while(run)
        {


            cam >> img; // получить кадр
            cout << "new image " << endl;

//
            Mat frame = img;
            Mat frameCopy = frame.clone();
            int frameWidth = frame.cols;
            int frameHeight = frame.rows;

            double t = (double) cv::getTickCount();
            Net net = readNetFromCaffe(protoFile, weightsFile);

            if (device == "cpu")
            {
                cout << "Using CPU device" << endl;
                net.setPreferableBackend(DNN_TARGET_CPU);
            }
            else if (device == "gpu")
            {
                cout << "Using GPU device" << endl;
                net.setPreferableBackend(DNN_BACKEND_CUDA);
                net.setPreferableTarget(DNN_TARGET_CUDA);
            }

            Mat inpBlob = blobFromImage(frame, 1.0 / 255, Size(inWidth, inHeight), Scalar(0, 0, 0), false, false);

            net.setInput(inpBlob);

            Mat output = net.forward();

            int H = output.size[2];
            int W = output.size[3];

            // find the position of the body parts
            vector<Point> points(nPoints);
            for (int n=0; n < nPoints; n++)
            {
                // Probability map of corresponding body's part.
                Mat probMap(H, W, CV_32F, output.ptr(0,n));

                Point2f p(-1,-1);
                Point maxLoc;
                double prob;
                minMaxLoc(probMap, 0, &prob, 0, &maxLoc);
                if (prob > thresh)
                {
                    p = maxLoc;
                    p.x *= (float)frameWidth / W ;
                    p.y *= (float)frameHeight / H ;

                    circle(frameCopy, cv::Point((int)p.x, (int)p.y), 8, Scalar(0,255,255), -1);
                    cv::putText(frameCopy, cv::format("%d", n), cv::Point((int)p.x, (int)p.y), cv::FONT_HERSHEY_COMPLEX, 1, cv::Scalar(0, 0, 255), 2);

                }
                points[n] = p;
            }

            int nPairs = sizeof(POSE_PAIRS)/sizeof(POSE_PAIRS[0]);
            serializeToJson(points);
            for (int n = 0; n < nPairs; n++)
            {
                // lookup 2 connected body/hand parts
                Point2f partA = points[POSE_PAIRS[n][0]];
                Point2f partB = points[POSE_PAIRS[n][1]];

                if (partA.x<=0 || partA.y<=0 || partB.x<=0 || partB.y<=0)
                    continue;

                line(frame, partA, partB, Scalar(0,255,255), 8);
                circle(frame, partA, 8, Scalar(0,0,255), -1);
                circle(frame, partB, 8, Scalar(0,0,255), -1);
            }

            t = ((double)cv::getTickCount() - t)/cv::getTickFrequency();
            cout << "Time Taken = " << t << endl;
            imshow(window, frame);
            /*imshow("Output-Keypoints", frameCopy);
            imshow("Output-Skeleton", frame);
            imwrite("Output-Skeleton.jpg", frame);*/
//
            run = (waitKey(33) != 27);
        }

        destroyAllWindows();
    }


    return 0;
}
