import os
import cv2
import numpy as np
from flask import Flask
from flask_restful import reqparse, abort, Api, Resource

app = Flask(__name__)
api = Api(app)

def createNewImg(orgFileName):
    # read image as grey scale
    fileName = os.path.splitext(orgFileName)[0]
    fileNameExtension = os.path.splitext(orgFileName)[1]
    Img = cv2.imread(orgFileName)
    
    # Scaling Bigger
    resizeBig = cv2.resize(src=Img, dsize=(Img.shape[1]*2, Img.shape[0]*2))
    fileNameResizeBig = fileName + '_scaling_big' + fileNameExtension
    cv2.imwrite(fileNameResizeBig, resizeBig)
    
    # Scaling Small
    resizeSmall = cv2.resize(Img,(64,64), interpolation=cv2.INTER_CUBIC)
    fileNameResizeSmall = fileName + '_scaling_small' + fileNameExtension
    cv2.imwrite(fileNameResizeSmall, resizeSmall)
    
    # x move 100px, y move 50px
    H = np.float32([[1,0,100],[0,1,50]])
    Translation = cv2.warpAffine(Img,H,dsize=(Img.shape[1], Img.shape[0])) 
    fileNameTranslation = fileName + '_translation' + fileNameExtension
    cv2.imwrite(fileNameTranslation, Translation)
    
    # Rotation left 90 degree
    RotateMatrix = cv2.getRotationMatrix2D(center=(Img.shape[1]/2, Img.shape[0]/2), angle=90, scale=1)
    Rotation = cv2.warpAffine(Img, RotateMatrix,dsize=(Img.shape[1], Img.shape[0]))
    fileNameRotation = fileName + '_rotation' + fileNameExtension
    cv2.imwrite(fileNameRotation, Rotation)
    
    # 3 point warpAffine
    # choose org image three point & choose new image three point
    pts1 = np.float32([[50,50],[200,50],[50,200]])
    pts2 = np.float32([[10,100],[200,50],[100,250]])
    M = cv2.getAffineTransform(pts1,pts2)
    WarpAffine = cv2.warpAffine(Img,M,dsize=(Img.shape[1]*2, Img.shape[0]*2))
    fileNamewarpAffine = fileName + '_warpAffine' + fileNameExtension
    cv2.imwrite(fileNamewarpAffine, WarpAffine)
    
    # 4 point warpPerspective
    rows,cols,ch = Img.shape
    pts1 = np.float32([[0,0],[0,cols],[rows,0],[rows,cols]])
    pts2 = np.float32([[0,50],[0,cols-50],[rows,-50],[rows,cols+50]])
    M = cv2.getPerspectiveTransform(pts1,pts2)
    Perspective = cv2.warpPerspective(Img,M,dsize=(Img.shape[1]*2, Img.shape[0]*2))
    fileNamewarpPerspective = fileName + '_Perspective' + fileNameExtension
    cv2.imwrite(fileNamewarpPerspective, Perspective)
    
parser = reqparse.RequestParser()
parser.add_argument('org_path')

class ConvertImg(Resource):
    def post(self):
        args = parser.parse_args()
        orgFileName = args['org_path']
        createNewImg(orgFileName)
        return 'test', 201

api.add_resource(ConvertImg, '/')

## Example cmd 
## curl http://localhost:5000/ -d "org_path=./org.jpg" -X POST -v
if __name__ == '__main__':
    app.run(debug=True)