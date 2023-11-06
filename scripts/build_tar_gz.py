"""Build a tar.gz file from a directory."""
import os
import tarfile


def build_tar_gz(tar_filename: str, directory: str) -> None:
    """Build a tar.gz file from a directory.

    :type tar_filename: str
    :param tar_filename: The name of the tar.gz file to create.
    :type directory: str
    :param directory: The directory to tar.gz.

    :rtype: None
    :return: None
    """
    with tarfile.open(tar_filename, 'w:gz') as tf:
        for dirname, subdirs, files in os.walk(directory):
            for filename in files:
                tf.add(os.path.join(dirname, filename), arcname=filename)


if __name__ == '__main__':
    base_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    wanted_dir = os.path.join(base_dir, "temp")
    linux_path = os.path.join(wanted_dir, "linux-amd64")
    darwin_path = os.path.join(wanted_dir, "darwin-amd64")
    build_tar_gz("linux-amd64.tar.gz", linux_path)
    build_tar_gz("darwin-amd64.tar.gz", darwin_path)
